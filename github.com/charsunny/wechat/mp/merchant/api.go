package merchant

import (
	"github.com/charsunny/wechat/mp/core"
)

// 创建商品.
func Create(clt *core.Client, product *Product) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/merchant/create?access_token="

	var result struct{
		core.Error
		ProductId string `json:"product_id"`
	}
	if err = clt.PostJSON(incompleteURL, product, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	product.ProductId = result.ProductId
	return
}

// 更新商品信息
func Update(clt *core.Client, product *Product) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/merchant/update?access_token="

	var result struct{
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, product, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}

	return
}

// 查询商品.
func Get(clt *core.Client, productId string) (product *Product, err error) {
	const incompleteURL = "https://api.weixin.qq.com/merchant/get?access_token="

	var request = struct {
		ProductId string `json:"product_id"`
	}{
		ProductId:productId,
	}

	var result struct {
		core.Error
		Product Product `json:"product_info"`
	}

	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	product = &result.Product
	return
}

// 获取指定状态的所有商品
// @param status 商品状态(0-全部, 1-上架, 2-下架)
func GetByStatus(clt *core.Client, status int) (list []*Product, err error)  {
	const incompleteURL = "https://api.weixin.qq.com/merchant/getbystatus?access_token="

	var request = struct {
		Status int `json:"status"`
	}{
		Status:status,
	}

	var result struct {
		core.Error
		Products []*Product `json:"products_info"`
	}

	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.Products
	return
}

// 商品上下架接口
// @param status
func UpdateStatus(clt *core.Client, productId string, status int) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/merchant/modproductstatus?access_token="

	var request = struct {
		ProductId string `json:"product_id"`
		Status int `json:"status"`
	}{
		ProductId:productId,
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

// 删除商品.
func Delete(clt *core.Client, productId string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/merchant/del?access_token="
	var request = struct {
		ProductId string `json:"product_id"`
	}{
		ProductId:productId,
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

// 获取分类的所有子分类接口
// @param cateId: 大分类ID(根节点分类id为1)
// 微信上所有的商品类目分为四级，估计是要拉四级类目
func GetSubCate(clt *core.Client, cateId string) (list []*CateInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/merchant/category/getsub?access_token="

	var request = struct {
		CateId string `json:"cate_id"`
	}{
		CateId:cateId,
	}

	var result struct {
		core.Error
		Cates []*CateInfo `json:"cate_list"`
	}

	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.Cates
	return
}

// 获取分类sku信息
func GetCateSku(clt *core.Client, cateId string) (list []*SkuTable, err error) {
	const incompleteURL = "https://api.weixin.qq.com/merchant/category/getsku?access_token="

	var request = struct {
		CateId string `json:"cate_id"`
	}{
		CateId:cateId,
	}

	var result struct {
		core.Error
		List []*SkuTable `json:"sku_table"`
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

// 获取分类属性信息
func GetCateProperties(clt *core.Client, cateId string) (list []*CateProperty, err error) {
	const incompleteURL = "https://api.weixin.qq.com/merchant/category/getproperty?access_token="

	var request = struct {
		CateId string `json:"cate_id"`
	}{
		CateId:cateId,
	}

	var result struct {
		core.Error
		List []*CateProperty `json:"properties"`
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