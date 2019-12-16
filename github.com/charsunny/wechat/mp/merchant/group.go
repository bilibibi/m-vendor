package merchant

import (
	"github.com/charsunny/wechat/mp/core"
)
// 分组管理接口

// 添加分组
func AddGroup(clt *core.Client, name string, productIds []string) (groupId int, err error) {
	const incompleteURL = "https://api.weixin.qq.com/merchant/group/add?access_token="
	type GroupDetail struct{
		GroupName string `json:"group_name"`
		ProductList []string `json:"product_list"`
	}
	var request = struct {
		detail *GroupDetail  `json:"group_detail"`
	}{
		detail: &GroupDetail{
			GroupName: name,
			ProductList: productIds,
		},
	}

	var result struct{
		core.Error
		GroupId int    `json:"group_id"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	groupId = result.GroupId
	return
}

// 编辑分组商品信息
func UpdateGroupProucts(clt *core.Client, groupId int, addList []string, removeList []string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/merchant/group/productmod?access_token="

	type ProductMod struct{
		ProductId string `json:"product_id"`
		ModAction int `json:"mod_action"`
	}

	var modlist []*ProductMod
	for _, prod := range addList {
		modlist = append(modlist, &ProductMod{
			ProductId:prod,
			ModAction:1,
		})
	}
	for _, prod := range removeList {
		modlist = append(modlist, &ProductMod{
			ProductId:prod,
			ModAction:0,
		})
	}

	var request = struct {
		GroupId int    `json:"group_id"`
		Products []*ProductMod `json:"product"`
	} {
		GroupId:groupId,
		Products:modlist,
	}

	var result struct{
		core.Error
		GroupId int    `json:"group_id"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	groupId = result.GroupId
	return
}

// 修改分组明成功
func UpdateGroup(clt *core.Client, groupId int, name string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/merchant/group/propertymod?access_token="

	var request = struct {
		GroupId int    `json:"group_id"`
		GroupName string `json:"group_name"`
	}{
		GroupId:groupId,
		GroupName:name,
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

// 删除分组信息
func DeleteGroup(clt *core.Client, groupId int) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/merchant/group/del?access_token="
	var request = struct {
		GroupId int    `json:"group_id"`
	}{
		GroupId:groupId,
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

// 获取邮费模板
// @param templateId 模板ID
func GetGroup(clt *core.Client,groupId int) (group *GroupInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/merchant/group/getbyid?access_token="

	var request = struct {
		GroupId int    `json:"group_id"`
	}{
		GroupId:groupId,
	}
	var result struct {
		core.Error
		Group *GroupInfo `json:"template_info"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	group = result.Group
	return
}

// 获取邮费模板列表
func GetGroupList(clt *core.Client) (list []*GroupInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/merchant/group/getall?access_token="


	var result struct {
		core.Error
		List []*GroupInfo `json:"groups_detail"`
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

