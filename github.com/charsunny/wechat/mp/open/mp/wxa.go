package mp

import (
	"github.com/charsunny/wechat/mp/core"
)

type WxaInfo struct {
	Username string `json:"username"`	// 小程序gh_id
	Status int `json:"status"`		//1：已关联； 2：等待小程序管理员确认中； 3：小程序管理员拒绝关联 12：等到公众号管理员确认中；
	Nickname string `json:"nickname"` // 小程序名称
	Selected int `json:"selected"`	// 是否在公众号管理页展示中
	NearbyDisplayStatus int `json:"nearby_display_status"`	// 是否展示在附近的小程序中
	Released int `json:"released"`	// 是否展示在附近的小程序中
	HeadimgUrl int `json:"headimg_url"`	// 是否展示在附近的小程序中
	FuncInfo []*FuncInfo `json:"func_info"`	// 是否展示在附近的小程序中
	Email string `json:"email"`	// 是否展示在附近的小程序中
}

type FuncInfo struct {
	Status int `json:"status"`
	Id int `json:"id"`
	Name string `json:"name"`
}

type AuthInfo struct {
	AppId string `json:"appid"`
	AuthorizationCode string `json:"authorization_code"`
	IsWxVerifySucc bool `json:"is_wx_verify_succ"`
	IsLinkSucc bool `json:"is_link_succ"`
}

// 获取公众号关联的小程序
func Get(clt *core.Client) (list []*WxaInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/wxamplinkget?access_token="

	var request = struct {
	}{
	}
	var result struct {
		core.Error
		WxOpens struct{
			Items []*WxaInfo `json:"items"`
		} `json:"wxopens"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.WxOpens.Items
	return
}

// 关联小程序
// 关联流程（需要公众号和小程序管理员双方确认）：
// 1.第三方平台调用接口发起关联
// 2 公众号管理员收到模板消息，同意关联小程序。
// 3.小程序管理员收到模板消息，同意关联公众号。
// 4.关联成功
// 等待管理员同意的中间状态可使用“获取公众号关联的小程序”接口进行查询。
// @param appid: 小程序appid
// @param notify_users: 是否发送模板消息通知公众号粉丝
// @param show_profile: 是否展示公众号主页中
func Bind(clt *core.Client, appid string, notify_users int, show_profile int) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/wxamplink?access_token="

	var request = struct {
		AppId string `json:"appid"`
		NotifyUsers int `json:"notify_users"`
		ShowProfile int `json:"show_profile"`
	}{
		AppId:appid,
		NotifyUsers:notify_users,
		ShowProfile:show_profile,
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

// 解除已关联的小程序
// @param appid: 小程序appid
func UnLink(clt *core.Client, appid string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/wxamplink?access_token="

	var request = struct {
		AppId string `json:"appid"`
	}{
		AppId:appid,
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

// 公众号快速注册小程序
// @param ticket 公众号管理员授权快速注册小程序返回的tickect，通过开放平台oauth的link获取
func FastRegister(clt *core.Client, ticket string) (info *AuthInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/account/fastregister?access_token="

	var request = struct {
		Ticket string `json:"ticket"`
	}{
		Ticket:ticket,
	}
	var result struct {
		core.Error
		AuthInfo
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.AuthInfo
	return
}