package merchant

import (
	"github.com/charsunny/wechat/mp/core"
)
// 库存管理相关接口

// 增加库存
// @param productId 商品ID
// @param skuInfo sku信息,格式"id1:vid1;id2:vid2",如商品为统一规格，则此处赋值为空字符串即可
// @param count 增加的库存数量
func AddStock(clt *core.Client, productId string, skuInfo string, count int) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/merchant/stock/add?access_token="

	var request = struct {
		ProductId string `json:"product_id"`
		SkuInfo string `json:"sku_info"`
		Count int `json:"quantity"`
	}{
		ProductId:productId,
		SkuInfo:skuInfo,
		Count:count,
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

// 减少库存
// @param productId 商品ID
// @param skuInfo sku信息,格式"id1:vid1;id2:vid2",如商品为统一规格，则此处赋值为空字符串即可
// @param count 增加的库存数量
func ReduceStock(clt *core.Client, productId string, skuInfo string, count int) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/merchant/stock/reduce?access_token="

	var request = struct {
		ProductId string `json:"product_id"`
		SkuInfo string `json:"sku_info"`
		Count int `json:"quantity"`
	}{
		ProductId:productId,
		SkuInfo:skuInfo,
		Count:count,
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