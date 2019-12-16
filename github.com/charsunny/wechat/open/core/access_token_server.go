package core

import (
	"net/url"
	"fmt"
	"time"
)


// DefaultAccessTokenServer 实现了 AccessTokenServer 接口.
//  NOTE:
//  1. 用于单进程环境.
//  2. 因为 DefaultAccessTokenServer 同时也是一个简单的中控服务器, 而不是仅仅实现 AccessTokenServer 接口,
//     所以整个系统只能存在一个 DefaultAccessTokenServer 实例!
type DefaultAccessTokenServer struct {

	appId      string

	client *Client

	refreshToken string

	token string // *accessToken

}

// NewDefaultAccessTokenServer 开放平台的DefaultAccessTokenServer, 用于获取开放平台授权appid的token
// 注意： appid为授权的公众号、小程序appid ， 必填字段
// token为授权token， 可以为空
// refreshtoken 用于刷新token的refresh token， 必填字段
// client 开放平台的client， 必填字段
func NewComponentAccessTokenServer(appId string, token string, refreshToken string, client *Client) (srv *DefaultAccessTokenServer) {

	if client == nil {
		panic("must init with a open platform client ")
	}

	if refreshToken == "" {
		panic("must init with a app refresh token ")
	}

	srv = &DefaultAccessTokenServer{
		appId:                    url.QueryEscape(appId),
		client:               client,
		token: token,
		refreshToken: refreshToken,
	}
	return
}

func (srv *DefaultAccessTokenServer) IID01332E16DF5011E5A9D5A4DB30FED8E1() {}

func (srv *DefaultAccessTokenServer) Token() (token string, err error) {
	cache := srv.client.AuthServer.CacheProvider()
	if cache != nil {
		token := cache.Get(srv.appId + "_access_token").(string)
		srv.token = token
	}
	if srv.token != "" {
		return srv.token, nil
	}
	return srv.RefreshToken("")
}


func (srv *DefaultAccessTokenServer) RefreshToken(currentToken string) (token string, err error) {

	url := "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token?component_access_token="
	params := map[string]interface{} {
		"component_appid": srv.client.AuthServer.AppId(),
		"authorizer_appid": srv.appId,
		"authorizer_refresh_token": srv.refreshToken,
	}
	var result struct{
		Error
		AppId string `json:"authorizer_appid"`
		AccessToken string `json:"authorizer_access_token"`
		ExpiresIn int `json:"expires_in"`
		RefreshToken string `json:"authorizer_refresh_token"`
	}

	srv.client.PostJSON(url, params, &result)

	if result.ErrCode != ErrCodeOK {
		err = &result.Error
		fmt.Println(err)
		return
	}
	cache := srv.client.AuthServer.CacheProvider()
	token = result.AccessToken
	if cache != nil {
		cache.Put(srv.appId + "_access_token", token, time.Duration(result.ExpiresIn * 1000 * 1000 * 1000))
	}
	srv.token = token
	return
}


type accessToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

