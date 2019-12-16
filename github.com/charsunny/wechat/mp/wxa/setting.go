package wxa

import (
	"fmt"
	"github.com/charsunny/wechat/mp/core"
)

// 设置小程序隐私设置（是否可被搜索）
func ChangeSearchStatus(clt *core.Client, status int) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/changewxasearchstatus?access_token="

	var request = struct {
		Status int `json:"status"`	//1表示不可搜索，0表示可搜索
	}{
		Status:status,
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

// 设置小程序隐私设置（是否可被搜索）
func QuerySearchStatus(clt *core.Client) (status int, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/changewxasearchstatus?access_token="

	var result struct {
		core.Error
		Status int `json:"status"`	//1表示不可搜索，0表示可搜索
	}
	if err = clt.GetJSON(incompleteURL,  &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	status = result.Status
	return
}

// 获取展示的公众号信息
func QueryShowWxaItem(clt *core.Client) (info* WxaItemInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/getshowwxaitem?access_token="

	var result struct {
		core.Error
		WxaItemInfo
	}
	if err = clt.GetJSON(incompleteURL,&result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.WxaItemInfo
	return
}

// 更新展示的公众号信息
func UpdateShowWxaItem(clt *core.Client, open int, appid string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/updateshowwxaitem?access_token="

	var request = struct {
		Open int `json:"wxa_subscribe_biz_flag"`	// 0 关闭，1 开启
		Appid string `json:"appid"`	// 如果开启，需要传新的公众号appid
	}{
		Open:open,
		Appid:appid,
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

// 获取可以用来设置的公众号列表
func GetShowWxaItemList(clt *core.Client, page, num int) (total int, list []*WxaItemInfo, err error) {
	incompleteURL := fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxamplinkforshow?page=%d&num=%d&access_token=", page, num)

	var result struct {
		core.Error
		TotalNum int `json:"total_num"`
		BizInfoList []*WxaItemInfo `json:"biz_info_list"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	total = result.TotalNum
	list = result.BizInfoList
	return
}