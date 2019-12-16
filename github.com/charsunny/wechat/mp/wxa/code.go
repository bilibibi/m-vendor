package wxa

import (
	"github.com/charsunny/wechat/mp/core"
	"fmt"
	"net/url"
)

// 绑定微信用户为小程序体验者
func CommitCode(clt *core.Client, template_id int, extjson string, version, desc string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/commit?access_token="

	var request = struct {
		TemplateId int `json:"template_id"`	//微信号
		ExtJson string `json:"ext_json"`	//微信号
		Version string `json:"user_version"`	//微信号
		Desc string `json:"user_desc"`	//微信号
	}{
		TemplateId:template_id,
		ExtJson:extjson,
		Version:version,
		Desc:desc,
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

// 获取小程序体验码
func GetExpQRCode(clt *core.Client, path string) (link string)  {
	incompleteURL := fmt.Sprintf("https://api.weixin.qq.com/wxa/get_qrcode?path=%s&access_token=", url.QueryEscape(path))

	if len(path) == 0 {
		incompleteURL = "https://api.weixin.qq.com/wxa/get_qrcode?access_token="
	}
	token, _ := clt.Token()
	link =  incompleteURL + token
	return
}

// 获取当前小程序类目
func GetCategory(clt *core.Client) (list []*WxaCategoryInfo, err error)  {
	const incompleteURL = "https://api.weixin.qq.com/wxa/get_category?access_token="

	var result struct {
		core.Error
		CategoryList []*WxaCategoryInfo `json:"category_list"`
	}
	if err = clt.GetJSON(incompleteURL,  &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.CategoryList
	return
}

// 获取小程序的第三方提交代码的页面配置
func GetPage(clt *core.Client) (list []string, err error)  {
	const incompleteURL = "https://api.weixin.qq.com/wxa/get_page?access_token="

	var result struct {
		core.Error
		PageList []string `json:"page_list"`
	}
	if err = clt.GetJSON(incompleteURL,  &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.PageList
	return
}

// 提交小程序审核
func SubmitCode(clt *core.Client, list []*WxaPageInfo, feedinfo, feedstuff string) (auditid int, err error)  {
	const incompleteURL = "https://api.weixin.qq.com/wxa/submit_audit?access_token="

	var request = struct {
		ItemList []*WxaPageInfo `json:"item_list"`	//微信号
		FeedbackInfo string `json:"feedback_info"`	// 反馈内容，不超过200字, 只有上个版本被驳回，才能使用“feedback_info”、“feedback_stuff”这两个字段，否则忽略处理
		FeedbackStuff string `json:"feedback_stuff"` //图片media_id列表，中间用“丨”分割，xx丨yy丨zz，不超过5张图片, 其中 media_id 可以通过新增临时素材接口上传而得到
	}{
		ItemList:list,
		FeedbackInfo:feedinfo,
		FeedbackStuff:feedstuff,
	}
	var result struct {
		core.Error
		Auditid int `json:"auditid"`	// 人员对应的唯一字符串
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	auditid = result.Auditid
	return
}


// 获取小程序审核状态
func GetAuditStatus(clt *core.Client, auditid int ) (status int, reason string, err error)  {
	const incompleteURL = "https://api.weixin.qq.com/wxa/get_auditstatus?access_token="

	var request = struct {
		Auditid int `json:"auditid"`	//微信号
	 }{
		Auditid:auditid,
	}
	var result struct {
		core.Error
		Status int `json:"status"`	// 审核状态，其中0为审核成功，1为审核失败，2为审核中，3已撤回
		Reason string `json:"reason"`	// 当status=1，审核被拒绝时，返回的拒绝原因
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	status = result.Status
	reason = result.Reason
	return
}

// 查询最新一次提交的审核状态
func GetLatestAuditStatus(clt *core.Client, auditid int ) (status int, reason string, err error)  {
	const incompleteURL = "https://api.weixin.qq.com/wxa/get_latest_auditstatus?access_token="


	var result struct {
		core.Error
		Status int `json:"status"`	// 审核状态，其中0为审核成功，1为审核失败，2为审核中，3已撤回
		Reason string `json:"reason"`	// 当status=1，审核被拒绝时，返回的拒绝原因
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	status = result.Status
	reason = result.Reason
	return
}

// 获取小程序审核状态
func ReleaseWxa(clt *core.Client ) (err error)  {
	const incompleteURL = "https://api.weixin.qq.com/wxa/release?access_token="

	var request = struct {

	}{

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


// 小程序版本回退
func RevertWxa(clt *core.Client ) (err error)  {
	const incompleteURL = "https://api.weixin.qq.com/wxa/revertcoderelease?access_token="

	var result struct {
		core.Error
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}

// 设置最低基础库版本
func SetWxSupportVersion(clt *core.Client, version string ) (err error)  {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/setweappsupportversion?access_token="

	var request = struct {
		Version string `json:"version"`
	}{
		Version:version,
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

// 小程序审核撤回, 单个帐号每天审核撤回次数最多不超过1次，一个月不超过10次
func UndoAudit(clt *core.Client ) (err error)  {
	const incompleteURL = "https://api.weixin.qq.com/wxa/undocodeaudit?access_token="

	var result struct {
		core.Error
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}

// TODO: 普通二维码跳转小程序 && 分阶段发布小程序