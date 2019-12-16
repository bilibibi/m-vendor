package core

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/charsunny/wechat/util/security"
	"github.com/charsunny/wechat/internal/debug/api"
	"github.com/charsunny/wechat/internal/debug/callback"
	iutil "github.com/charsunny/wechat/internal/util"
	"github.com/charsunny/wechat/util"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
	"unicode"
	"unsafe"
)

type Cache interface {
	Get(key string) interface{}
	GetMulti(keys []string) []interface{}
	Put(key string, val interface{}, timeout time.Duration) error
	Delete(key string) error
	Incr(key string) error
	Decr(key string) error
	IsExist(key string) bool
	ClearAll() error
	StartAndGC(config string) error
}

// AuthServer 用于处理微信服务器的回调请求, 并发安全!
//  通常情况下一个 AuthServer 实例用于处理一个开放平台的回调消息, 并且刷新平台的token和tickprovider;
type AuthServer struct {
	appId string

	appSecret string

	tokenBucketPtrMutex sync.Mutex     // used only by writers
	tokenBucketPtr      unsafe.Pointer // *tokenBucket

	aesKeyBucketPtrMutex sync.Mutex     // used only by writers
	aesKeyBucketPtr      unsafe.Pointer // *aesKeyBucket

	componentVerifyTicketPtrMutex sync.Mutex     // used only by writers
	componentVerifyTicketPtr      unsafe.Pointer // *componentVerifyTicket

	httpClient *http.Client

	refreshTokenRequestChan  chan string             // chan currentToken
	refreshTokenResponseChan chan refreshTokenResult // chan {token, err}

	tokenCache unsafe.Pointer // *accessToken

	errorHandler ErrorHandler

	cacheProvider Cache
}


func (srv *AuthServer) AppId() string {
	return srv.appId
}

func (srv *AuthServer) AppSecret() string {
	return srv.appSecret
}

func (srv *AuthServer) CacheProvider() Cache {
	return srv.cacheProvider
}


type tokenBucket struct {
	currentToken string
	lastToken    string
}

type aesKeyBucket struct {
	currentAESKey []byte
	lastAESKey    []byte
}

type componentTicketBucket struct {
	currentTikect string
	lastTicket   string
}

// NewAuthServer 创建一个新的 AuthServer.
//  appId:        必选; 开放平台的AppId;
//  appSecret:    必选; 开放平台的AppSecret;
//  token:        必须; 开放平台用于验证签名的token;
//  base64AESKey: 必选; 开放平台处理ticket回掉的aeskey;
//  handler:      必须; 处理微信服务器推送过来的消息(事件)的Handler;
//  cacheProvider: 用于缓存provider ticker的缓存， ticker 10分钟刷新一次， 如果服务器crash或者重启，10分钟内授权可能出问题.
func NewAuthServer(appId, appSecret, token, base64AESKey string, httpClient *http.Client, cacheProvider Cache) (srv *AuthServer) {
	if token == "" {
		panic("empty token")
	}

	if httpClient == nil {
		httpClient = util.DefaultHttpClient
	}


	var (
		aesKey []byte
		err    error
	)
	if base64AESKey != "" {
		if len(base64AESKey) != 43 {
			panic("the length of base64AESKey must equal to 43")
		}
		aesKey, err = base64.StdEncoding.DecodeString(base64AESKey + "=")
		if err != nil {
			panic(fmt.Sprintf("Decode base64AESKey:%s failed", base64AESKey))
		}
	}

	srv = &AuthServer{
		appId:           appId,
		appSecret: appSecret,
		tokenBucketPtr:  unsafe.Pointer(&tokenBucket{currentToken: token}),
		aesKeyBucketPtr: unsafe.Pointer(&aesKeyBucket{currentAESKey: aesKey}),
		componentVerifyTicketPtr: unsafe.Pointer(&componentTicketBucket{currentTikect: ""}),
		errorHandler:    DefaultErrorHandler,
		httpClient: httpClient,
		cacheProvider: cacheProvider,
		refreshTokenRequestChan:  make(chan string),
		refreshTokenResponseChan: make(chan refreshTokenResult),
	}

	go srv.tokenUpdateDaemon(time.Hour * 100)

	// 首先从cache中读取上一次的保存的ticker provider， 不必从微信服务端获取
	if srv.cacheProvider != nil {
		fmt.Printf("get cache provider: %v", srv.cacheProvider)
		if  ticker, ok := srv.cacheProvider.Get("component_ticker").(string); ok {
			fmt.Printf("get coo ticker: %s", ticker)
			srv.setComponentVerifyTicket(ticker)
		}
	}

	return
}

func (srv *AuthServer) getComponentVerifyTicket() (currentTicket, lastTicket string) {
	if p := (*componentTicketBucket)(atomic.LoadPointer(&srv.componentVerifyTicketPtr)); p != nil {
		return p.currentTikect, p.lastTicket
	}
	return
}

// SetToken 设置签名token.
func (srv *AuthServer) setComponentVerifyTicket(ticket string) (err error) {
	if ticket == "" {
		return errors.New("empty ticket")
	}

	srv.componentVerifyTicketPtrMutex.Lock()
	defer srv.componentVerifyTicketPtrMutex.Unlock()

	currentTicket, lastTicket := srv.getComponentVerifyTicket()

	if ticket == currentTicket {
		return
	}

	bucket := componentTicketBucket{
		currentTikect: ticket,
		lastTicket:    currentTicket,
	}
	atomic.StorePointer(&srv.componentVerifyTicketPtr, unsafe.Pointer(&bucket))

	if lastTicket == "" {	// get componet token when last ticket is nil
		srv.ComponentToken()
	}
	return
}

func (srv *AuthServer) removeLastComponentVerifyTicket(lastToken string) {
	srv.componentVerifyTicketPtrMutex.Lock()
	defer srv.componentVerifyTicketPtrMutex.Unlock()

	currentTicket2, lastTicket2 := srv.getComponentVerifyTicket()
	if currentTicket2 != lastTicket2 {
		return
	}

	bucket := componentTicketBucket{
		currentTikect: lastTicket2,
	}
	atomic.StorePointer(&srv.componentVerifyTicketPtr, unsafe.Pointer(&bucket))
	return
}

func (srv *AuthServer) getToken() (currentToken, lastToken string) {
	if p := (*tokenBucket)(atomic.LoadPointer(&srv.tokenBucketPtr)); p != nil {
		return p.currentToken, p.lastToken
	}
	return
}

// SetToken 设置签名token.
func (srv *AuthServer) SetToken(token string) (err error) {
	if token == "" {
		return errors.New("empty token")
	}

	srv.tokenBucketPtrMutex.Lock()
	defer srv.tokenBucketPtrMutex.Unlock()

	currentToken, _ := srv.getToken()
	if token == currentToken {
		return
	}

	bucket := tokenBucket{
		currentToken: token,
		lastToken:    currentToken,
	}
	atomic.StorePointer(&srv.tokenBucketPtr, unsafe.Pointer(&bucket))
	return
}

func (srv *AuthServer) removeLastToken(lastToken string) {
	srv.tokenBucketPtrMutex.Lock()
	defer srv.tokenBucketPtrMutex.Unlock()

	currentToken2, lastToken2 := srv.getToken()
	if lastToken != lastToken2 {
		return
	}

	bucket := tokenBucket{
		currentToken: currentToken2,
	}
	atomic.StorePointer(&srv.tokenBucketPtr, unsafe.Pointer(&bucket))
	return
}

func (srv *AuthServer) getAESKey() (currentAESKey, lastAESKey []byte) {
	if p := (*aesKeyBucket)(atomic.LoadPointer(&srv.aesKeyBucketPtr)); p != nil {
		return p.currentAESKey, p.lastAESKey
	}
	return
}

// SetAESKey 设置aes加密解密key.
//  base64AESKey: aes加密解密key, 43字节长(base64编码, 去掉了尾部的'=').
func (srv *AuthServer) SetAESKey(base64AESKey string) (err error) {
	if len(base64AESKey) != 43 {
		return errors.New("the length of base64AESKey must equal to 43")
	}
	aesKey, err := base64.StdEncoding.DecodeString(base64AESKey + "=")
	if err != nil {
		return
	}

	srv.aesKeyBucketPtrMutex.Lock()
	defer srv.aesKeyBucketPtrMutex.Unlock()

	currentAESKey, _ := srv.getAESKey()
	if bytes.Equal(aesKey, currentAESKey) {
		return
	}

	bucket := aesKeyBucket{
		currentAESKey: aesKey,
		lastAESKey:    currentAESKey,
	}
	atomic.StorePointer(&srv.aesKeyBucketPtr, unsafe.Pointer(&bucket))
	return
}

func (srv *AuthServer) removeLastAESKey(lastAESKey []byte) {
	srv.aesKeyBucketPtrMutex.Lock()
	defer srv.aesKeyBucketPtrMutex.Unlock()

	currentAESKey2, lastAESKey2 := srv.getAESKey()
	if !bytes.Equal(lastAESKey, lastAESKey2) {
		return
	}

	bucket := aesKeyBucket{
		currentAESKey: currentAESKey2,
	}
	atomic.StorePointer(&srv.aesKeyBucketPtr, unsafe.Pointer(&bucket))
	return
}

func (srv *AuthServer) ComponentToken() (token string, err error) {
	if p := (*componentAccessToken)(atomic.LoadPointer(&srv.tokenCache)); p != nil {
		return p.Token, nil
	}
	return srv.ComponentRefreshToken("")
}

type refreshTokenResult struct {
	token string
	err   error
}

func (srv *AuthServer) ComponentRefreshToken(currentToken string) (token string, err error) {
	srv.refreshTokenRequestChan <- currentToken
	rslt := <-srv.refreshTokenResponseChan
	return rslt.token, rslt.err
}

func (srv *AuthServer) tokenUpdateDaemon(initTickDuration time.Duration) {
	tickDuration := initTickDuration

NEW_TICK_DURATION:
	ticker := time.NewTicker(tickDuration)

	for {
		select {
		case currentToken := <-srv.refreshTokenRequestChan:
			accessToken, cached, err := srv.updateToken(currentToken)
			fmt.Printf("%v,%v,%v", accessToken, cached, err)
			if err != nil {
				srv.refreshTokenResponseChan <- refreshTokenResult{err: err}
				break
			}
			srv.refreshTokenResponseChan <- refreshTokenResult{token: accessToken.Token}
			if !cached {
				tickDuration = time.Duration(accessToken.ExpiresIn) * time.Second
				ticker.Stop()
				goto NEW_TICK_DURATION
			}

		case <-ticker.C:
			accessToken, _, err := srv.updateToken("")
			if err != nil {
				break
			}
			newTickDuration := time.Duration(accessToken.ExpiresIn) * time.Second
			if abs(tickDuration-newTickDuration) > time.Second*5 {
				tickDuration = newTickDuration
				ticker.Stop()
				goto NEW_TICK_DURATION
			}
		}
	}
}

func abs(x time.Duration) time.Duration {
	if x >= 0 {
		return x
	}
	return -x
}

type componentAccessToken struct {
	Token     string `json:"component_access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

// updateToken 从微信服务器获取新的 access_token 并存入缓存, 同时返回该 access_token.
func (srv *AuthServer) updateToken(currentToken string) (token *componentAccessToken, cached bool, err error) {

	if currentToken != "" {
		if p := (*componentAccessToken)(atomic.LoadPointer(&srv.tokenCache)); p != nil && currentToken != p.Token {
			return p, true, nil // 无需更改 p.ExpiresIn 参数值, cached == true 时用不到
		}
	}

	ticket, lasttikect :=  srv.getComponentVerifyTicket()
	if ticket == "" && lasttikect == "" {
		atomic.StorePointer(&srv.tokenCache, nil)
		err = fmt.Errorf("ticket empty. Server ticket is empty, get ticket latter\n")
		return
	}
	url := "https://api.weixin.qq.com/cgi-bin/component/api_component_token"
	fmt.Println(ticket, srv.appSecret, srv.appId)
	req, _ := json.Marshal(map[string]string{
		"component_appid": srv.AppId(),
		"component_appsecret": srv.AppSecret(),
		"component_verify_ticket": ticket,
	})
	httpResp, err := srv.httpClient.Post(url, "application/json; charset=utf-8", bytes.NewReader(req))
	if err != nil {
		if lasttikect != "" {
			req, _ = json.Marshal(map[string]string{
				"component_appid": srv.AppId(),
				"component_appsecret": srv.AppSecret(),
				"component_verify_ticket": lasttikect,
			})
			httpResp, err = srv.httpClient.Post(url, "application/json; charset=utf-8", bytes.NewReader(req))
			if err != nil {
				srv.removeLastComponentVerifyTicket(ticket)
				atomic.StorePointer(&srv.tokenCache, nil)
				return
			}
		} else {
			srv.removeLastComponentVerifyTicket(ticket)
			atomic.StorePointer(&srv.tokenCache, nil)
			return
		}

	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		atomic.StorePointer(&srv.tokenCache, nil)
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	var result struct {
		Error
		componentAccessToken
	}
	if err = api.DecodeJSONHttpResponse(httpResp.Body, &result); err != nil {
		atomic.StorePointer(&srv.tokenCache, nil)
		return
	}
	if result.ErrCode != ErrCodeOK {
		atomic.StorePointer(&srv.tokenCache, nil)
		err = &result.Error
		return
	}

	// 由于网络的延时, access_token 过期时间留有一个缓冲区
	switch {
	case result.ExpiresIn > 31556952: // 60*60*24*365.2425
		atomic.StorePointer(&srv.tokenCache, nil)
		err = errors.New("expires_in too large: " + strconv.FormatInt(result.ExpiresIn, 10))
		return
	case result.ExpiresIn > 60*60:
		result.ExpiresIn -= 60 * 10
	case result.ExpiresIn > 60*30:
		result.ExpiresIn -= 60 * 5
	case result.ExpiresIn > 60*5:
		result.ExpiresIn -= 60
	case result.ExpiresIn > 60:
		result.ExpiresIn -= 10
	default:
		atomic.StorePointer(&srv.tokenCache, nil)
		err = errors.New("expires_in too small: " + strconv.FormatInt(result.ExpiresIn, 10))
		return
	}

	tokenCopy := result.componentAccessToken
	atomic.StorePointer(&srv.tokenCache, unsafe.Pointer(&tokenCopy))
	token = &tokenCopy
	return
}

// ServeHTTP 处理微信服务器的回调请求, query 参数可以为 nil.
func (srv *AuthServer) ServeHTTP(w http.ResponseWriter, r *http.Request, query url.Values) (msg []byte) {

	callback.DebugPrintRequest(r)
	if query == nil {
		query = r.URL.Query()
	}
	errorHandler := srv.errorHandler
	switch r.Method {
	case "POST": // 推送消息(事件)
		switch encryptType := query.Get("encrypt_type"); encryptType {
		case "aes":
			haveSignature := query.Get("signature")
			if haveSignature == "" {
				errorHandler.ServeError(w, r, errors.New("not found signature query parameter"))
				return
			}
			haveMsgSignature := query.Get("msg_signature")
			if haveMsgSignature == "" {
				errorHandler.ServeError(w, r, errors.New("not found msg_signature query parameter"))
				return
			}
			timestampString := query.Get("timestamp")
			if timestampString == "" {
				errorHandler.ServeError(w, r, errors.New("not found timestamp query parameter"))
				return
			}
			_, err := strconv.ParseInt(timestampString, 10, 64)
			if err != nil {
				err = fmt.Errorf("can not parse timestamp query parameter %q to int64", timestampString)
				errorHandler.ServeError(w, r, err)
				return
			}
			nonce := query.Get("nonce")
			if nonce == "" {
				errorHandler.ServeError(w, r, errors.New("not found nonce query parameter"))
				return
			}

			var token string
			currentToken, lastToken := srv.getToken()
			if currentToken == "" {
				err = errors.New("token was not set for AuthServer, see NewAuthServer function or AuthServer.SetToken method")
				errorHandler.ServeError(w, r, err)
				return
			}
			token = currentToken
			wantSignature := iutil.Sign(token, timestampString, nonce)
			if !security.SecureCompareString(haveSignature, wantSignature) {
				if lastToken == "" {
					err = fmt.Errorf("check signature failed, have: %s, want: %s", haveSignature, wantSignature)
					errorHandler.ServeError(w, r, err)
					return
				}
				token = lastToken
				wantSignature = iutil.Sign(token, timestampString, nonce)
				if !security.SecureCompareString(haveSignature, wantSignature) {
					err = fmt.Errorf("check signature failed, have: %s, want: %s", haveSignature, wantSignature)
					errorHandler.ServeError(w, r, err)
					return
				}
			} else {
				if lastToken != "" {
					srv.removeLastToken(lastToken)
				}
			}

			buffer := textBufferPool.Get().(*bytes.Buffer)
			buffer.Reset()
			defer textBufferPool.Put(buffer)

			if _, err = buffer.ReadFrom(r.Body); err != nil {
				errorHandler.ServeError(w, r, err)
				return
			}
			requestBodyBytes := buffer.Bytes()

			var requestHttpBody cipherRequestHttpBody
			if err = xmlUnmarshal(requestBodyBytes, &requestHttpBody); err != nil {
				errorHandler.ServeError(w, r, err)
				return
			}

			wantMsgSignature := iutil.MsgSign(token, timestampString, nonce, string(requestHttpBody.Base64EncryptedMsg))
			if !security.SecureCompareString(haveMsgSignature, wantMsgSignature) {
				err = fmt.Errorf("check msg_signature failed, have: %s, want: %s", haveMsgSignature, wantMsgSignature)
				errorHandler.ServeError(w, r, err)
				return
			}

			encryptedMsg := make([]byte, base64.StdEncoding.DecodedLen(len(requestHttpBody.Base64EncryptedMsg)))
			encryptedMsgLen, err := base64.StdEncoding.Decode(encryptedMsg, requestHttpBody.Base64EncryptedMsg)
			if err != nil {
				errorHandler.ServeError(w, r, err)
				return
			}
			encryptedMsg = encryptedMsg[:encryptedMsgLen]

			var aesKey []byte
			currentAESKey, lastAESKey := srv.getAESKey()
			if currentAESKey == nil {
				err = errors.New("aes key was not set for AuthServer, see NewAuthServer function or AuthServer.SetAESKey method")
				errorHandler.ServeError(w, r, err)
				return
			}
			aesKey = currentAESKey
			_, msgPlaintext, haveAppIdBytes, err := iutil.AESDecryptMsg(encryptedMsg, aesKey)
			if err != nil {
				if lastAESKey == nil {
					errorHandler.ServeError(w, r, err)
					return
				}
				aesKey = lastAESKey
				_, msgPlaintext, haveAppIdBytes, err = iutil.AESDecryptMsg(encryptedMsg, aesKey)
				if err != nil {
					errorHandler.ServeError(w, r, err)
					return
				}
			} else {
				if lastAESKey != nil {
					srv.removeLastAESKey(lastAESKey)
				}
			}
			callback.DebugPrintPlainRequestMessage(msgPlaintext)

			haveAppId := string(haveAppIdBytes)
			wantAppId := srv.appId
			if wantAppId != "" && !security.SecureCompareString(haveAppId, wantAppId) {
				err = fmt.Errorf("the message AppId mismatch, have: %s, want: %s", haveAppId, wantAppId)
				errorHandler.ServeError(w, r, err)
				return
			}

			var verifyTicket verifyTickectMsg
			if err = xml.Unmarshal(msgPlaintext, &verifyTicket); err != nil {
				errorHandler.ServeError(w, r, err)
				return
			}
			if wantAppId != verifyTicket.AppId {
				err = fmt.Errorf("the message AppId mismatch between ciphertext and plaintext, %q != %q",
					wantAppId, verifyTicket.AppId)
				errorHandler.ServeError(w, r, err)
				return
			}
			if verifyTicket.InfoType == "notify_third_fasteregister" {
				msg = msgPlaintext
				return
			}
			srv.setComponentVerifyTicket(verifyTicket.ComponentVerifyTicket)
			// set cookie for later user
			// 首先从cache中读取上一次的保存的ticker provider， 不必从微信服务端获取
			if srv.cacheProvider != nil {	// 做10分钟的缓存，用于服务器恢复缓存
				srv.cacheProvider.Delete("component_ticker")	// 删除老的 添加新的
				srv.cacheProvider.Put("component_ticker", verifyTicket.ComponentVerifyTicket, time.Duration(time.Second * 60 * 60))	 // 缓存1小时，方便调用
			}
			fmt.Printf("get wechat server ticker: %v, %v, %v\n", wantAppId, verifyTicket.ComponentVerifyTicket, w.Header())
			io.WriteString(w, "success")
			return msgPlaintext
		default:
			errorHandler.ServeError(w, r, errors.New("unknown encrypt_type: "+encryptType))
		}
	}
	io.WriteString(w, "success")
	return
}

// =====================================================================================================================

type verifyTickectMsg struct {
	XMLName            struct{} `xml:"xml"`
	AppId string `xml:"AppId"`
	CreateTime string `xml:"CreateTime"`
	InfoType string `xml:"InfoType"`
	ComponentVerifyTicket string `xml:"ComponentVerifyTicket"`

}

type cipherRequestHttpBody struct {
	XMLName            struct{} `xml:"xml"`
	AppId         string   `xml:"AppId"`
	Base64EncryptedMsg []byte   `xml:"Encrypt"`
}

var (
	msgStartElementLiteral = []byte("<xml>")
	msgEndElementLiteral   = []byte("</xml>")

	msgAppIdStartElementLiteral = []byte("<AppId>")
	msgAppIdEndElementLiteral   = []byte("</AppId>")

	msgEncryptStartElementLiteral = []byte("<Encrypt>")
	msgEncryptEndElementLiteral   = []byte("</Encrypt>")

	cdataStartLiteral = []byte("<![CDATA[")
	cdataEndLiteral   = []byte("]]>")
)

//<xml>
//    <AppId><![CDATA[gh_b1eb3f8bd6c6]]></AppId>
//    <Encrypt><![CDATA[DlCGq+lWQuyjNNK+vDaO0zUltpdUW3u4V00WCzsdNzmZGEhrU7TPxG52viOKCWYPwTMbCzgbCtakZHyNxr5hjoZJ7ORAUYoIAGQy/LDWtAnYgDO+ppKLp0rDq+67Dv3yt+vatMQTh99NII6x9SEGpY3O2h8RpG99+NYevQiOLVKqiQYzan21sX/jE4Y3wZaeudsb4QVjqzRAPaCJ5nS3T31uIR9fjSRgHTDRDOzjQ1cHchge+t6faUhniN5VQVTE+wIYtmnejc55BmHYPfBnTkYah9+cTYnI3diUPJRRiyVocJyHlb+XOZN22dsx9yzKHBAyagaoDIV8Yyb/PahcUbsqGv5wziOgLJQIa6z93/VY7d2Kq2C2oBS+Qb+FI9jLhgc3RvCi+Yno2X3cWoqbsRwoovYdyg6jme/H7nMZn77PSxOGRt/dYiWx2NuBAF7fNFigmbRiive3DyOumNCMvA==]]></Encrypt>
//</xml>
func xmlUnmarshal(data []byte, p *cipherRequestHttpBody) error {
	data = bytes.TrimSpace(data)
	if !bytes.HasPrefix(data, msgStartElementLiteral) || !bytes.HasSuffix(data, msgEndElementLiteral) {
		log.Printf("[WARNING] xmlUnmarshal failed, data:\n%s\n", data)
		return xml.Unmarshal(data, p)
	}
	data2 := data[len(msgStartElementLiteral) : len(data)-len(msgEndElementLiteral)]

	// AppId
	AppIdElementBytes := data2
	i := bytes.Index(AppIdElementBytes, msgAppIdStartElementLiteral)
	if i == -1 {
		log.Printf("[WARNING] xmlUnmarshal failed, data:\n%s\n", data)
		return xml.Unmarshal(data, p)
	}
	AppIdElementBytes = AppIdElementBytes[i+len(msgAppIdStartElementLiteral):]
	AppIdElementBytes = bytes.TrimLeftFunc(AppIdElementBytes, unicode.IsSpace)
	if !bytes.HasPrefix(AppIdElementBytes, cdataStartLiteral) {
		log.Printf("[WARNING] xmlUnmarshal failed, data:\n%s\n", data)
		return xml.Unmarshal(data, p)
	}
	AppIdElementBytes = AppIdElementBytes[len(cdataStartLiteral):]
	i = bytes.Index(AppIdElementBytes, cdataEndLiteral)
	if i == -1 {
		log.Printf("[WARNING] xmlUnmarshal failed, data:\n%s\n", data)
		return xml.Unmarshal(data, p)
	}
	AppId := AppIdElementBytes[:i]
	AppIdElementBytes = AppIdElementBytes[i+len(cdataEndLiteral):]
	AppIdElementBytes = bytes.TrimLeftFunc(AppIdElementBytes, unicode.IsSpace)
	if !bytes.HasPrefix(AppIdElementBytes, msgAppIdEndElementLiteral) {
		log.Printf("[WARNING] xmlUnmarshal failed, data:\n%s\n", data)
		return xml.Unmarshal(data, p)
	}
	AppIdElementBytes = AppIdElementBytes[len(msgAppIdEndElementLiteral):]

	// Encrypt
	EncryptElementBytes := AppIdElementBytes
	i = bytes.Index(EncryptElementBytes, msgEncryptStartElementLiteral)
	if i == -1 {
		EncryptElementBytes = data2
		i = bytes.Index(EncryptElementBytes, msgEncryptStartElementLiteral)
		if i == -1 {
			log.Printf("[WARNING] xmlUnmarshal failed, data:\n%s\n", data)
			return xml.Unmarshal(data, p)
		}
	}
	EncryptElementBytes = EncryptElementBytes[i+len(msgEncryptStartElementLiteral):]
	EncryptElementBytes = bytes.TrimLeftFunc(EncryptElementBytes, unicode.IsSpace)
	if !bytes.HasPrefix(EncryptElementBytes, cdataStartLiteral) {
		log.Printf("[WARNING] xmlUnmarshal failed, data:\n%s\n", data)
		return xml.Unmarshal(data, p)
	}
	EncryptElementBytes = EncryptElementBytes[len(cdataStartLiteral):]
	i = bytes.Index(EncryptElementBytes, cdataEndLiteral)
	if i == -1 {
		log.Printf("[WARNING] xmlUnmarshal failed, data:\n%s\n", data)
		return xml.Unmarshal(data, p)
	}
	Encrypt := EncryptElementBytes[:i]
	EncryptElementBytes = EncryptElementBytes[i+len(cdataEndLiteral):]
	EncryptElementBytes = bytes.TrimLeftFunc(EncryptElementBytes, unicode.IsSpace)
	if !bytes.HasPrefix(EncryptElementBytes, msgEncryptEndElementLiteral) {
		log.Printf("[WARNING] xmlUnmarshal failed, data:\n%s\n", data)
		return xml.Unmarshal(data, p)
	}

	p.AppId = string(AppId)
	p.Base64EncryptedMsg = Encrypt
	return nil
}
