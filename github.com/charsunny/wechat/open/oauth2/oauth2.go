package oauth2

import (
	mpoauth2 "github.com/charsunny/wechat/mp/oauth2"
	"net/http"
	"net/url"
	"strconv"
)

type AuthorizationInfo struct {
	AppId string `json:"authorizer_appid"`
	AccessToken string `json:"authorizer_access_token"`
	ExpiresIn int `json:"expires_in"`
	RefreshToken string `json:"authorizer_refresh_token"`
	FuncInfo [] interface{} `json:"func_info"`

}

// AuthWebURL 生成网页连接的开放平台授权地址.
//  appId:       开放平台
// 	preAuthCode: 预授权code， 从平台获取
// authType : 1则商户点击链接后，手机端仅展示公众号、2表示仅展示小程序，3表示公众号和小程序都展示。如果为未指定，则默认小程序和公众号都展示
//  redirectURI: 授权后重定向的回调链接地址
func OpenAuthWebURL(appId, redirectURI, preAuthCode string, authType int) string {
	if authType == 0 {	// 如果auth type 是0 强制变成3 不然会出现授权错误
		authType = 3
	}
	return "https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=" + url.QueryEscape(appId) +
		"&redirect_uri=" + url.QueryEscape(redirectURI) +
		"&pre_auth_code=" + preAuthCode +
		"&auth_type=" + strconv.Itoa(authType)
}

// OpenAuthWechatLink 生成微信内点击的开放平台授权地址.
//  appId:       开放平台
// 	preAuthCode: 预授权code， 从平台获取
// 	authType : 1则商户点击链接后，手机端仅展示公众号、2表示仅展示小程序，3表示公众号和小程序都展示。如果为未指定，则默认小程序和公众号都展示
//  redirectURI: 授权后重定向的回调链接地址
func OpenAuthWechatLink(appId, redirectURI, preAuthCode string, authType int) string {
	if authType == 0 {	// 如果auth type 是0 强制变成3 不然会出现授权错误
		authType = 3
	}
	return "https://mp.weixin.qq.com/safe/bindcomponent?action=bindcomponent&no_scan=1&component_appid=" + url.QueryEscape(appId) +
		"&redirect_uri=" + url.QueryEscape(redirectURI) +
		"&pre_auth_code=" + preAuthCode +
		"&auth_type=" + strconv.Itoa(authType) +
		"#wechat_redirect"
}

// WebAuthCodeURL 生成网页授权地址.
//  appId:       网页的唯一标识
//  redirectURI: 授权后重定向的回调链接地址
//  scope:       应用授权作用域
//  state:       重定向后会带上 state 参数, 开发者可以填写 a-zA-Z0-9 的参数值, 最多128字节
func WebAuthCodeURL(appId, redirectURI, scope, state string) string {
	return "https://open.weixin.qq.com/connect/qrconnect?appid=" + url.QueryEscape(appId) +
		"&redirect_uri=" + url.QueryEscape(redirectURI) +
		"&response_type=code&scope=" + url.QueryEscape(scope) +
		"&state=" + url.QueryEscape(state) +
		"#wechat_redirect"
}

// FastRegisterWxaLink 快速注册小程序的开放平台连接.
// appId:       公众号的 appid
// component_appid: 第三方平台的appid
// redirectURI: 授权后重定向的回调链接地址
// 公众号管理员扫码后在手机端完成授权确认。
// 跳转回第三方平台，会在上述 redirect_uri后拼接 ticket=*
func FastRegisterWxaLink(appId, componentAppid, redirectURI string) string {
	return "https://mp.weixin.qq.com/cgi-bin/fastregisterauth?copy_wx_verify=1&component_appid=" + url.QueryEscape(componentAppid)  +
		"&appid=" + url.QueryEscape(appId) + "&redirect_uri=" + url.QueryEscape(redirectURI)
}

// Auth 检验授权凭证 access_token 是否有效.
//  accessToken: 网页授权接口调用凭证
//  openId:      用户的唯一标识
//  httpClient:  如果不指定则默认为 util.DefaultHttpClient
func WebAuth(accessToken, openId string, httpClient *http.Client) (valid bool, err error) {
	return mpoauth2.Auth(accessToken, openId, httpClient)
}


