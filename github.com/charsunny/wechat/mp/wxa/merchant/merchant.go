package merchant

import (
	"github.com/charsunny/wechat/mp/core"
)

// 绑定微信用户为小程序体验者
func GetDistrict(clt *core.Client) (list [][]map[string]interface{}, version string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/get_district?access_token="

	var result struct {
		core.Error
		Status int `json:"status"`
		Version string `json:"data_version"`
		Result [][]map[string]interface{} `json:"result"`
	}

	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	version = result.Version
	list = result.Result
	return
}
