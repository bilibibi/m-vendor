package open

import (
	"github.com/charsunny/wechat/mp/core"
)

// 该API用于创建一个开放平台帐号，并将一个尚未绑定开放平台帐号的公众号/小程序绑定至该开放平台帐号上。
// 新创建的开放平台帐号的主体信息将设置为与之绑定的公众号或小程序的主体。
// @param appId: 授权公众号或小程序的 appid
func Create(clt *core.Client, appId string) (openAppId string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/open/create?access_token="

	var request = struct {
		Appid string `json:"appid"`
	}{
		Appid: appId,
	}
	var result struct {
		core.Error
		OpenAppId string `json:"open_appid"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	openAppId = result.OpenAppId
	return
}


// 该API用于将一个尚未绑定开放平台帐号的公众号或小程序绑定至指定开放平台帐号上。二者须主体相同。
// @param openAppId: 开放平台帐号appid
// @param appId: 授权公众号或小程序的 appid
func Bind(clt *core.Client, appId string, openAppId string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/open/bind?access_token="

	var request = struct {
		Appid string `json:"appid"`
		OpenAppId string `json:"open_appid"`
	}{
		Appid: appId,
		OpenAppId:openAppId,
	}
	var result struct {
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}

// 该API用于将一个尚未绑定开放平台帐号的公众号或小程序绑定至指定开放平台帐号上。二者须主体相同。
// @param openAppId: 开放平台帐号appid
// @param appId: 授权公众号或小程序的 appid
func Unbind(clt *core.Client, appId string, openAppId string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/open/unbind?access_token="

	var request = struct {
		Appid string `json:"appid"`
		OpenAppId string `json:"open_appid"`
	}{
		Appid: appId,
		OpenAppId:openAppId,
	}
	var result struct {
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}

// 该API用于获取公众号或小程序所绑定的开放平台帐号。
// @param appId: 授权公众号或小程序的 appid
func Get(clt *core.Client, appId string) (openAppId string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/open/get?access_token="

	var request = struct {
		Appid string `json:"appid"`
	}{
		Appid: appId,
	}
	var result struct {
		core.Error
		OpenAppId string `json:"open_appid"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	openAppId = result.OpenAppId
	return
}