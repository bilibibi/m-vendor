package template

import (
	"github.com/charsunny/wechat/mp/core"
)

type Keyword struct {
	KeywordId string `json:"keyword_id"`
	Name string `json:"name"`
	Example string `json:"example"`
}

type Template struct {
	Id string `json:"id"`		// 作为模板时返回
	Title string `json:"title"` // 作为模板时和作为添加到模板库后都返回
	TemplateId string `json:"template_id"` // 作为添加到模板库后返回
	Content string `json:"content"` // 作为添加到模板库后返回
	Example string `json:"example"` // 作为添加到模板库后返回
}

type WxaTemplateMsg struct {
	TemplateId string `json:"template_id"`	// 模板id
	Page string `json:"page"`	//跳转小程序页面
	FormId string `json:"form_id"`	// formid
	Data map[string]interface{} `json:"data"`	// 模板数据
	EmphasisKeyword string `json:"emphasis_keyword"`	//加重数据
}

type MpTemplateMsg struct {
	TemplateId string `json:"template_id"`	// 模板id
	AppId string `json:"appid"`	//跳转小程序页面
	Url string `json:"url"`	// formid
	Data map[string]interface{} `json:"data"`	// 模板数据
	MiniProgram struct{	//跳转小程序页面
		AppId string `json:"appid"`
		PagePath string `json:"pagepath"`
	} `json:"miniprogram"`
}

// 组合模板并添加至帐号下的个人模板库
// id	string		是	模板标题id，可通过接口获取，也可登录小程序后台查看获取
// keyword_id_list	[]int 是  开发者自行组合好的模板关键词列表，关键词顺序可以自由搭配（例如[3,5,4]或[4,5,3]），最多支持10个关键词组合
func AddTemplate(clt *core.Client, id string, keyIds []int) (template_id string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/template/add?access_token="

	var request = struct {
		Id string `json:"id"`	//路径
		KeywordIdList []int `json:"keyword_id_list"`	//宽度
	}{
		Id:id,
		KeywordIdList:keyIds,
	}
	var result struct {
		core.Error
		TemplateId string `json:"template_id"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	template_id = result.TemplateId
	return
}

// 组合模板并添加至帐号下的个人模板库
// template_id	string		要删除的模板id
func DeleteTemplate(clt *core.Client, template_id string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/template/del?access_token="

	var request = struct {
		TemplateId string `json:"template_id"`	//路径
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

// 获取模板的keyword列表
// id	string		是	模板标题id，可通过接口获取，也可登录小程序后台查看获取
func GetTemplateLibraryById(clt *core.Client, id string) (title string, keywords[]*Keyword, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/template/library/get?access_token="

	var request = struct {
		Id string `json:"id"`	//路径
	}{
		Id:id,
	}
	var result struct {
		core.Error
		Title string `json:"title"`
		KeywordList []*Keyword `json:"keyword_list"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	title = result.Title
	keywords = result.KeywordList
	return
}
// 获取模板的keyword列表
// id	string		是	模板标题id，可通过接口获取，也可登录小程序后台查看获取
func GetTemplateLibraryList(clt *core.Client, page, count int) (total_count int, list[]*Template, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/template/library/list?access_token="

	var request = struct {
		Offset int `json:"offset"`	//路径
		Count int `json:"count"`	//路径
	}{
		Offset: page,
		Count: count,
	}
	var result struct {
		core.Error
		TotalCount int `json:"total_count"`
		List []*Template `json:"list"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	total_count = result.TotalCount
	list = result.List
	return
}

// 获取小程序模板库列表
// count: 用于分页，表示拉取count条记录。最大为 20。最后一页的list长度可能小于请求的count。
func GetTemplateList(clt *core.Client, page, count int) (list[]*Template, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/template/list?access_token="

	var request = struct {
		Offset int `json:"offset"`	//路径
		Count int `json:"count"`	//路径
	}{
		Offset: page,
		Count: count,
	}
	var result struct {
		core.Error
		List []*Template `json:"list"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.List
	return
}

// 发送模板消息
// data: { "k1":{"value":""}}
// emphasis_keyword:  "keyword1.DATA"
func Send(clt *core.Client, touser, template_id, page, form_id string, data map[string]interface{}, emphasis_keyword string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token="

	var request = struct {
		ToUser string `json:"touser"`	// 用户openid
		TemplateId string `json:"template_id"`	// 模板id
		Page string `json:"page"`	//跳转小程序页面
		FormId string `json:"form_id"`	// formid
		Data map[string]interface{} `json:"data"`	// 模板数据
		EmphasisKeyword string `json:"emphasis_keyword"`	//加重数据
	}{
		ToUser: touser,
		TemplateId: template_id,
		Page: page,
		FormId: form_id,
		Data: data,
		EmphasisKeyword: emphasis_keyword,
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

// 统一发送接口
func UniformSend(clt *core.Client, touser string, msg *WxaTemplateMsg, mpmsg *MpTemplateMsg) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send?access_token="

	var request = struct {
		ToUser string `json:"touser"`	// 用户openid
		WxaMsg *WxaTemplateMsg `json:"weapp_template_msg"`
		MpMsg *MpTemplateMsg `json:"mp_template_msg"`
	}{
		ToUser: touser,
		WxaMsg:msg,
		MpMsg:mpmsg,
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
