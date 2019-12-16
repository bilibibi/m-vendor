package wxa

import (
	"fmt"
	"net/url"
	"github.com/charsunny/wechat/mp/core"
)

// 获取支付的unionid信息
func GetPaidUnionId(clt *core.Client, openid string, transId string) (unionid string, err error) {
	incompleteURL := fmt.Sprintf("https://api.weixin.qq.com/wxa/getpaidunionid?openid=%s&transaction_id=%s&access_token=",  url.QueryEscape(openid), transId)

	var result struct {
		core.Error
		UnionId string `json:"unionid"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	unionid = result.UnionId
	return
}

// 获取帐号基本信息
func GetBaseInfo(clt *core.Client) (info *WxaInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/account/getaccountbasicinfo?access_token="

	var result struct {
		core.Error
		WxaInfo
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.WxaInfo
	return
}

// 小程序改名
func SetNickName(clt *core.Client, info *WxaNameRequestInfo) (wording string, audit_id int, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/setnickname?access_token="

	var result struct {
		core.Error
		Wording string `json:"wording"`	// 材料说明
		AuditId int `json:"audit_id"`	// 审核单id 若接口未返回audit_id，说明名称已直接设置成功，无需审核；若返回audit_id则名称正在审核中
	}
	if err = clt.PostJSON(incompleteURL, info, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	wording = result.Wording
	audit_id = result.AuditId
	return
}

// 小程序改名
func QuerySetNickStatus(clt *core.Client,  audit_id int) (info *WxaNameResultInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/api_wxa_querynickname?access_token="

	var request = struct {
		AuditId int `json:"audit_id"`
	}{
		AuditId:audit_id,
	}
	var result struct {
		core.Error
		WxaNameResultInfo
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.WxaNameResultInfo
	return
}

// 微信认证名称检测
func CheckNickName(clt *core.Client,  nickname string) (wording string, hit bool, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxverify/checkwxverifynickname?access_token="

	var request = struct {
		NickName string `json:"nick_name"`
	}{
		NickName:nickname,
	}
	var result struct {
		core.Error
		Wording string `json:"wording"`	// 命中关键字的说明描述（给用户看的）
		HitCondition bool `json:"hit_condition"`	// 是否命中关键字策略。若命中，可以选填关键字材料
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	wording = result.Wording
	hit = result.HitCondition
	return
}

// 修改小程序头像
func ModifyHeadImage(clt *core.Client,  head_img_media_id string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/account/modifyheadimage?access_token="

	var request = struct {
		HeadImgMediaId string `json:"head_img_media_id"`
		X1 float64 `json:"x1"`
		X2 float64 `json:"x2"`
		Y1 float64 `json:"y1"`
		Y2 float64 `json:"y2"`
	}{
		HeadImgMediaId:head_img_media_id,
		X1: 0,
		X2: 1,
		Y1: 0,
		Y2: 1,
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

// 修改小程序功能介绍
func ModifySignature(clt *core.Client,  signature string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/account/modifysignature?access_token="

	var request = struct {
		Signature string `json:"signature"`
	}{
		Signature:signature,
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

// TODO:  换绑小程序管理员接口

// 获取所有可设置类目接口
func GetAllCategories(clt *core.Client) (list []*WxaCategory, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/getallcategories?access_token="

	var result struct {
		core.Error
		List struct{
			List []*WxaCategory `json:"categories"`
		} `json:"categories_list"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.List.List
	return
}

// 添加类目
func AddCategory(clt *core.Client,  list [] *WxaActionCateInfo) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/addcategory?access_token="

	var request = struct {
		 List [] *WxaActionCateInfo `json:"categories"`
	}{
		List:list,
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

// 删除类目
func DeleteCategory(clt *core.Client,  first, second int) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/deletecategory?access_token="

	var request = struct {
		First int `json:"first"`
		Second int `json:"second"`
	}{
		First:first,
		Second:second,
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

// 获取商户的设置类目
func GetAccountCategory(clt *core.Client) (item *CategoryItem, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/getcategory?access_token="

	var result struct {
		core.Error
		*CategoryItem
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	item = result.CategoryItem
	return
}

// 添加类目
func ModifyCategory(clt *core.Client,  item *WxaActionCateInfo) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/modifycategory?access_token="

	var result struct {
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, item, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}

// 修改域名
func ModifyDomain(clt *core.Client,  action string, requestdomain , wsrequestdomain, uploaddoamin, downloaddomain []string) ( requst, ws, upload, download []string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/modify_domain?access_token="

	var req = struct{
		Action string `json:"action"`
		RequestDomain []string `json:"requestdomain"`
		WsrequestDomain []string `json:"wsrequestdomain"`
		UploadDomain []string `json:"uploaddomain"`
		DownloadDomain []string `json:"downloaddomain"`
	} {
		Action:action,
		RequestDomain:requestdomain,
		WsrequestDomain:wsrequestdomain,
		UploadDomain:uploaddoamin,
		DownloadDomain:downloaddomain,
	}
	var result struct {
		core.Error
		RequestDomain []string `json:"requestdomain"`
		WsrequestDomain []string `json:"wsrequestdomain"`
		UploadDomain []string `json:"uploaddomain"`
		DownloadDomain []string `json:"downloaddomain"`
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	requst = result.RequestDomain
	ws = result.WsrequestDomain
	upload = result.UploadDomain
	download = result.DownloadDomain
	return
}

// 设置webviewdomain
func SetWebviewDomain(clt *core.Client, action string, domains []string) (list []string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/setwebviewdomain?access_token="

	var req = struct{
		Action string `json:"action"`
		Domains []string `json:"webviewdomain"`
	} {
		Action:action,
		Domains: domains,
	}
	var result struct {
		core.Error
		Domains []string `json:"webviewdomain"`
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.Domains
	return
}