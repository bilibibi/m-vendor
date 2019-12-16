package wxa

import (
	"github.com/charsunny/wechat/open/core"
)

type WxaCodeDraftInfo struct {
	DraftId int `json:"draft_id"`	// 展示的公众号appid
	UserDesc string `json:"user_desc"`	// 展示的公众号nickname
	UserVersion string `json:"user_version"`	// 展示的公众号头像
	CreateTime int `json:"create_time"`	// 是否可以设置 1 可以，0，不可以
}

type WxaCodeTemplateInfo struct {
	TemplateId int `json:"template_id"`	// 展示的公众号appid
	UserDesc string `json:"user_desc"`	// 展示的公众号nickname
	UserVersion string `json:"user_version"`	// 展示的公众号头像
	CreateTime int `json:"create_time"`	// 是否可以设置 1 可以，0，不可以
}

type FastRegesiterInfo struct {
	Name string `json:"name"`	// 企业名（需与工商部门登记信息一致）
	Code string `json:"code"`	// 证件编号
	CodeType int `json:"code_type"`	// 企业代码类型（1：统一社会信用代码， 2：组织机构代码，3：营业执照注册号）
	LegalPersonaWechat string `json:"legal_persona_wechat"`	 // 法人微信
	LegalPersonaName string `json:"legal_persona_name"`	 // 法人姓名
	ComponentPhone string `json:"component_phone"`	 //第三方联系电话
}

// 快速注册小程序
func FastRegisterWeapp(clt *core.Client, info *FastRegesiterInfo) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/component/fastregisterweapp?action=create&component_access_token="

	var result struct {
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, info, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}

// 查询快速注册小程序状态
func QueryFastRegisterWeapp(clt *core.Client, info *FastRegesiterInfo) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/component/fastregisterweapp?action=search&component_access_token="
	var req = struct{
		Name string `json:"name"`
		LegalPersonaWechat string `json:"legal_persona_wechat"`	 // 法人微信
		LegalPersonaName string `json:"legal_persona_name"`	 // 法人姓名
	} {
		 Name: info.Name,
		 LegalPersonaName: info.LegalPersonaName,
		 LegalPersonaWechat: info.LegalPersonaWechat,
	}
	var result struct {
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, &req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}

// 绑定微信用户为小程序体验者
func GetCodeTemplateDraftList(clt *core.Client) (list []*WxaCodeDraftInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/gettemplatedraftlist?access_token="

	var result struct {
		core.Error
		TemplateList []*WxaCodeDraftInfo `json:"draft_list"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.TemplateList
	return
}

// 获取代码模版库中的所有小程序代码模版
func GetCodeTemplateList(clt *core.Client) (list []*WxaCodeTemplateInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/gettemplatelist?access_token="

	var result struct {
		core.Error
		TemplateList []*WxaCodeTemplateInfo `json:"template_list"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.TemplateList
	return
}


// 将草稿箱的草稿选为小程序代码模版
func AddToTemplate(clt *core.Client, draft_id int) (err error)  {
	const incompleteURL = "https://api.weixin.qq.com/wxa/addtotemplate?access_token="

	var request = struct {
		DraftId int `json:"draft_id"`	// 反馈内容，不超过200字, 只有上个版本被驳回，才能使用“feedback_info”、“feedback_stuff”这两个字段，否则忽略处理
	}{
		DraftId:draft_id,
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

// 将草稿箱的草稿选为小程序代码模版
func DeleteTemplate(clt *core.Client, template_id int) (err error)  {
	const incompleteURL = "https://api.weixin.qq.com/wxa/deletetemplate?access_token="

	var request = struct {
		TemplateId int `json:"template_id"`	// 反馈内容，不超过200字, 只有上个版本被驳回，才能使用“feedback_info”、“feedback_stuff”这两个字段，否则忽略处理
	}{
		TemplateId:template_id,
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
