package merchant

import (
	"github.com/charsunny/wechat/mp/core"
)
// 邮费模板管理接口

// 增加邮费模板
func AddExpress(clt *core.Client, express *Express) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/merchant/express/add?access_token="

	var result struct{
		core.Error
		TemplateId int    `json:"template_id"`
	}
	if err = clt.PostJSON(incompleteURL, express, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	express.TemplateId = result.TemplateId
	return
}

// 更新邮费模板
func UpdateExpress(clt *core.Client, express *Express) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/merchant/express/update?access_token="

	var request = struct {
		TemplateId int    `json:"template_id"`
		Template *Express `json:"delivery_template"`
	}{
		TemplateId:express.TemplateId,
		Template:express,
	}

	var result struct{
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

// 删除邮费模板
func DeleteExpress(clt *core.Client, templateId int) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/merchant/express/del?access_token="
	var request = struct {
		TemplateId int    `json:"template_id"`
	}{
		TemplateId:templateId,
	}
	var result core.Error
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 获取邮费模板详情
// @param templateId 模板ID
func GetExpress(clt *core.Client,templateId int) (express *Express, err error) {
	const incompleteURL = "https://api.weixin.qq.com/merchant/express/getbyid?access_token="

	var request = struct {
		TemplateId int    `json:"template_id"`
	}{
		TemplateId:templateId,
	}
	var result struct {
		core.Error
		Express *Express `json:"template_info"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	express = result.Express
	return
}

// 获取邮费模板列表
func GetExpressList(clt *core.Client) (list []*Express, err error) {
	const incompleteURL = "https://api.weixin.qq.com/merchant/express/getall?access_token="


	var result struct {
		core.Error
		List []*Express `json:"templates_info"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.List
	return
}

