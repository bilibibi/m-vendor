package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/charsunny/wechat/internal/debug/api"
	"github.com/charsunny/wechat/internal/debug/api/retry"
	"github.com/charsunny/wechat/util"
	"io/ioutil"
)

type Client struct {
	AccessTokenServer
	HttpClient *http.Client
}

// NewClient 创建一个新的 Client.
//  如果 clt == nil 则默认用 util.DefaultHttpClient
func NewClient(srv AccessTokenServer, clt *http.Client) *Client {
	if srv == nil {
		panic("nil AccessTokenServer")
	}
	if clt == nil {
		clt = util.DefaultHttpClient
	}
	return &Client{
		AccessTokenServer: srv,
		HttpClient:        clt,
	}
}

// GetJSON HTTP GET 微信资源, 然后将微信服务器返回的 JSON 用 encoding/json 解析到 response.
//
//  NOTE:
//  1. 一般不需要调用这个方法, 请直接调用高层次的封装函数;
//  2. 最终的 URL == incompleteURL + access_token;
//  3. response 格式有要求, 要么是 *Error, 要么是下面结构体的指针(注意 Error 必须是第一个 Field):
//      struct {
//          Error
//          ...
//      }
func (clt *Client) GetJSON(incompleteURL string, response interface{}) (err error) {
	ErrorStructValue, ErrorErrCodeValue := checkResponse(response)

	httpClient := clt.HttpClient
	if httpClient == nil {
		httpClient = util.DefaultHttpClient
	}

	token, err := clt.Token()
	if err != nil {
		return
	}

	hasRetried := false
RETRY:
	finalURL := incompleteURL + url.QueryEscape(token)
	if err = httpGetJSON(httpClient, finalURL, response); err != nil {
		return
	}

	switch errCode := ErrorErrCodeValue.Int(); errCode {
	case ErrCodeOK:
		return
	case ErrCodeInvalidCredential, ErrCodeAccessTokenExpired:
		errMsg := ErrorStructValue.Field(errorErrMsgIndex).String()
		retry.DebugPrintError(errCode, errMsg, token)
		if !hasRetried {
			hasRetried = true
			ErrorStructValue.Set(errorZeroValue)
			if token, err = clt.RefreshToken(token); err != nil {
				return
			}
			retry.DebugPrintNewToken(token)
			goto RETRY
		}
		retry.DebugPrintFallthrough(token)
		fallthrough
	default:
		return
	}
}

func httpGetJSON(clt *http.Client, url string, response interface{}) error {
	api.DebugPrintGetRequest(url)
	httpResp, err := clt.Get(url)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	s, _ := ioutil.ReadAll(httpResp.Body) //把  body 内容读入字符串 s
	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}
	fmt.Printf("resp body: %s \n", s)
	err = json.Unmarshal(s, response)
	if err != nil {
		fmt.Printf("json decode err: %v\n", err)
	}
	return err
}

// PostJSON 用 encoding/json 把 request marshal 为 JSON, HTTP POST 到微信服务器,
// 然后将微信服务器返回的 JSON 用 encoding/json 解析到 response.
//
//  NOTE:
//  1. 一般不需要调用这个方法, 请直接调用高层次的封装函数;
//  2. 最终的 URL == incompleteURL + access_token;
//  3. response 格式有要求, 要么是 *Error, 要么是下面结构体的指针(注意 Error 必须是第一个 Field):
//      struct {
//          Error
//          ...
//      }
func (clt *Client) PostJSON(incompleteURL string, request interface{}, response interface{}) (err error) {
	ErrorStructValue, ErrorErrCodeValue := checkResponse(response)

	buffer := textBufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer textBufferPool.Put(buffer)

	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	if err = encoder.Encode(request); err != nil {
		return
	}
	requestBodyBytes := buffer.Bytes()
	if i := len(requestBodyBytes) - 1; i >= 0 && requestBodyBytes[i] == '\n' {
		requestBodyBytes = requestBodyBytes[:i] // 去掉最后的 '\n', 这样能统一log格式, 不然可能多一个空白行
	}
	httpClient := clt.HttpClient
	if httpClient == nil {
		httpClient = util.DefaultHttpClient
	}

	token, err := clt.Token()
	if err != nil {
		return
	}

	hasRetried := false
RETRY:
	finalURL := incompleteURL + url.QueryEscape(token)
	if _, err = httpPostJSON(httpClient, finalURL, requestBodyBytes, response); err != nil {
		return
	}
	fmt.Println(err)
	switch errCode := ErrorErrCodeValue.Int(); errCode {
	case ErrCodeOK:
		return
	case ErrCodeInvalidCredential, ErrCodeAccessTokenExpired:
		errMsg := ErrorStructValue.Field(errorErrMsgIndex).String()
		retry.DebugPrintError(errCode, errMsg, token)
		if !hasRetried {
			hasRetried = true
			ErrorStructValue.Set(errorZeroValue)
			if token, err = clt.RefreshToken(token); err != nil {
				return
			}
			retry.DebugPrintNewToken(token)
			goto RETRY
		}
		retry.DebugPrintFallthrough(token)
		fallthrough
	default:
		return
	}
}

// PostData  HTTP POST 到微信服务器,
// 然后将微信服务器返回的 JSON 用 encoding/json 解析到 response.
//
//  NOTE:
//  1. 一般不需要调用这个方法, 请直接调用高层次的封装函数;
//  2. 最终的 URL == incompleteURL + access_token;
//  3. response 格式有要求, 要么是 *Error, 要么是下面结构体的指针(注意 Error 必须是第一个 Field):
//      struct {
//          Error
//          ...
//      }
func (clt *Client) PostData(incompleteURL string, requestBodyBytes []byte, response interface{}) (err error) {
	ErrorStructValue, ErrorErrCodeValue := checkResponse(response)

	httpClient := clt.HttpClient
	if httpClient == nil {
		httpClient = util.DefaultHttpClient
	}

	token, err := clt.Token()
	if err != nil {
		return
	}

	hasRetried := false
RETRY:
	finalURL := incompleteURL + url.QueryEscape(token)
	if _, err = httpPostJSON(httpClient, finalURL, requestBodyBytes, response); err != nil {
		fmt.Println(err)
		return
	}

	switch errCode := ErrorErrCodeValue.Int(); errCode {
	case ErrCodeOK:
		return
	case ErrCodeInvalidCredential, ErrCodeAccessTokenExpired:
		errMsg := ErrorStructValue.Field(errorErrMsgIndex).String()
		retry.DebugPrintError(errCode, errMsg, token)
		if !hasRetried {
			hasRetried = true
			ErrorStructValue.Set(errorZeroValue)
			if token, err = clt.RefreshToken(token); err != nil {
				return
			}
			retry.DebugPrintNewToken(token)
			goto RETRY
		}
		retry.DebugPrintFallthrough(token)
		fallthrough
	default:
		return
	}
}

// PostData  HTTP POST 到微信服务器,
// 然后将微信服务器返回的 JSON 用 encoding/json 解析到 response.
//
//  NOTE:
//  1. 一般不需要调用这个方法, 请直接调用高层次的封装函数;
//  2. 最终的 URL == incompleteURL + access_token;
//  3. response 格式有要求, 要么是 *Error, 要么是下面结构体的指针(注意 Error 必须是第一个 Field):
//      struct {
//          Error
//          ...
//      }
func (clt *Client) PostJsonData(incompleteURL string, request interface{}, response interface{}) (data []byte, err error) {
	ErrorStructValue, ErrorErrCodeValue := checkResponse(response)

	buffer := textBufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer textBufferPool.Put(buffer)

	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	if err = encoder.Encode(request); err != nil {
		return
	}
	requestBodyBytes := buffer.Bytes()
	if i := len(requestBodyBytes) - 1; i >= 0 && requestBodyBytes[i] == '\n' {
		requestBodyBytes = requestBodyBytes[:i] // 去掉最后的 '\n', 这样能统一log格式, 不然可能多一个空白行
	}
	httpClient := clt.HttpClient
	if httpClient == nil {
		httpClient = util.DefaultHttpClient
	}

	token, err := clt.Token()
	if err != nil {
		return
	}

	hasRetried := false
RETRY:
	finalURL := incompleteURL + url.QueryEscape(token)
	if data, err = httpPostJSON(httpClient, finalURL, requestBodyBytes, response); err != nil {
		return
	}
	fmt.Println(err)
	switch errCode := ErrorErrCodeValue.Int(); errCode {
	case ErrCodeOK:
		return
	case ErrCodeInvalidCredential, ErrCodeAccessTokenExpired:
		errMsg := ErrorStructValue.Field(errorErrMsgIndex).String()
		retry.DebugPrintError(errCode, errMsg, token)
		if !hasRetried {
			hasRetried = true
			ErrorStructValue.Set(errorZeroValue)
			if token, err = clt.RefreshToken(token); err != nil {
				return
			}
			retry.DebugPrintNewToken(token)
			goto RETRY
		}
		retry.DebugPrintFallthrough(token)
		fallthrough
	default:
		return
	}
}


func httpPostJSON(clt *http.Client, url string, body []byte, response interface{}) (data []byte, err error) {
	api.DebugPrintPostJSONRequest(url, body)
	httpResp, err := clt.Post(url, "application/json; charset=utf-8", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()
	data, _ = ioutil.ReadAll(httpResp.Body) //把  body 内容读入字符串 s
	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http.Status: %s", httpResp.Status)
	}
	err = json.Unmarshal(data, response)
	if err != nil {
		fmt.Printf("json decode err: %v\n", err)
		return data, err
	}
	fmt.Printf("resp body: %s \n", response)
	return data, err
}

// checkResponse 检查 response 参数是否满足特定的结构要求, 如果不满足要求则会 panic, 否则返回相应的 reflect.Value.
func checkResponse(response interface{}) (ErrorStructValue, ErrorErrCodeValue reflect.Value) {
	responseValue := reflect.ValueOf(response)
	if responseValue.Kind() != reflect.Ptr {
		panic("the type of response is incorrect")
	}
	responseStructValue := responseValue.Elem()
	if responseStructValue.Kind() != reflect.Struct {
		panic("the type of response is incorrect")
	}

	if t := responseStructValue.Type(); t == errorType {
		ErrorStructValue = responseStructValue
	} else {
		if t.NumField() == 0 {
			panic("the type of response is incorrect")
		}
		v := responseStructValue.Field(0)
		if v.Type() == errorType {
			ErrorStructValue = v
		} else {
			panic("the type of response is incorrect")
		}
	}
	ErrorErrCodeValue = ErrorStructValue.Field(errorErrCodeIndex)
	return
}
