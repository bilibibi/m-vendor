package wxa

import (
	"github.com/charsunny/wechat/mp/core"
)

// 绑定微信用户为小程序体验者
func BindTester(clt *core.Client, wechatid string) (userstr string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/bind_tester?access_token="

	var request = struct {
		Wechatid string `json:"wechatid"`	//微信号
	}{
		Wechatid:wechatid,
	}
	var result struct {
		core.Error
		Userstr string `json:"userstr"`	// 人员对应的唯一字符串
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	userstr = result.Userstr
	return
}

// 解除绑定小程序的体验者
func UnbindTester(clt *core.Client, wechatid string, userstr string) (err error)  {
	const incompleteURL = "https://api.weixin.qq.com/wxa/unbind_tester?access_token="

	var request = struct {
		Wechatid string `json:"wechatid"`	//微信号
		Userstr string `json:"userstr"`	// 人员对应的唯一字符串（可通过获取体验者api获取已绑定人员的字符串，userstr和wechatid填写其中一个即可）
	}{
		Wechatid:wechatid,
		Userstr:userstr,
	}
	var result struct {
		core.Error
		Userstr string `json:"userstr"`	// 人员对应的唯一字符串
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

// 获取体验者列表
func GetBindTester(clt *core.Client) (list []string, err error)  {
	const incompleteURL = "https://api.weixin.qq.com/wxa/memberauth?access_token="

	var request = struct {
		Action string `json:"action"`	//微信号
	}{
		Action:"get_experiencer",
	}
	var result struct {
		core.Error
		Members []struct{
			UserStr string `json:"userstr"`
		} `json:"members"`	// 人员对应的唯一字符串
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	for _, str := range result.Members {
		list = append(list, str.UserStr)
	}
	return
}

