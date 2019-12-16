package qrcode

import (
	"github.com/charsunny/wechat/mp/core"
)

// 接口生成的小程序码，永久有效，有数量限制
// path	string		是	扫码进入的小程序页面路径，
// width	number	430	否	二维码的宽度，单位 px。最小 280px，最大 1280px
func CreateQRCode(clt *core.Client, path string, width int) (data []byte, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token="

	var request = struct {
		Path string `json:"path"`	//路径
		Width int `json:"width"`	//宽度
	}{
		Path:path,
		Width:width,
	}
	var result struct {
		core.Error
	}
	data, err = clt.PostJsonData(incompleteURL, &request, &result)
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}

type LineColor struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

// 接口生成的小程序码，永久有效，有数量限制
// path	string		是	扫码进入的小程序页面路径，必须是已经发布的小程序存在的页面（否则报错），例如 pages/index/index, 根路径前不要填加 /,不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面
// width	number	430	否	二维码的宽度，单位 px。最小 280px，最大 1280px
// auto_color	boolean	false	否	自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调，默认 false
// line_color	Object	{"r":0,"g":0,"b":0}	否	auto_color 为 false 时生效，
// is_hyaline	boolean	false	否	是否需要透明底色，为 true 时，生成透明底色的小程序
func CreateWxaCode(clt *core.Client, path string, width int, lineColor *LineColor, auto_color, is_hyaline bool) (data []byte, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/getwxacode?access_token="

	var request = struct {
		Path string `json:"path"`	//路径
		Width int `json:"width"`	//宽度
		AutoColor bool `json:"auto_color"`
		LineColor *LineColor `json:"line_color,omitempty"`
		IsHyaline bool `json:"is_hyaline"`
	}{
		Path:path,
		Width:width,
		AutoColor:auto_color,
		LineColor:lineColor,
		IsHyaline:is_hyaline,
	}
	var result struct {
		core.Error
	}
	data, err = clt.PostJsonData(incompleteURL, &request, &result)
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}

// 接口生成的小程序码，永久有效，有数量限制
// path	string		是	扫码进入的小程序页面路径，必须是已经发布的小程序存在的页面（否则报错），例如 pages/index/index, 根路径前不要填加 /,不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面
// width	number	430	否	二维码的宽度，单位 px。最小 280px，最大 1280px
// auto_color	boolean	false	否	自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调，默认 false
// line_color	Object	{"r":0,"g":0,"b":0}	否	auto_color 为 false 时生效，
// is_hyaline	boolean	false	否	是否需要透明底色，为 true 时，生成透明底色的小程序
func CreateWxaCodeUnlimited(clt *core.Client, scene string, path string, width int, lineColor *LineColor, auto_color, is_hyaline bool) (data []byte, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token="

	var request = struct {
		Path string `json:"page"`	//路径
		Width int `json:"width"`	//宽度
		Scene string `json:"scene"`
		AutoColor bool `json:"auto_color"`
		LineColor *LineColor `json:"line_color,omitempty"`
		IsHyaline bool `json:"is_hyaline"`
	}{
		Path:path,
		Width:width,
		AutoColor:auto_color,
		LineColor:lineColor,
		Scene: scene,
		IsHyaline:is_hyaline,
	}
	var result struct {
		core.Error
	}
	data, err = clt.PostJsonData(incompleteURL, &request, &result)

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	err = nil
	return
}

type JumpRule struct {
	Prefix string `json:"prefix"` // 二维码规则
	PermitSubRule int `json:"permit_sub_rule"`	// 是否独占符合二维码前缀匹配规则的所有子规 1 为不占用，2 为占用; 详见
	Path string `json:"path"`	// 小程序功能页面
	OpenVersion int `json:"open_version"`	// 1 开发版 2 体验版 3 正式版
	DebugUrl []string `json:"debug_url,omitempty"`	// 测试链接（选填）可填写不多于 5 个用于测试的二维码完整链接
	State int `json:"state"` 	//发布标志位，1 表示未发布，2 表示已发布
	IsEdit int `json:"is_edit"`	// 编辑标志位，0 表示新增二维码规则，1 表示修改已有二维码规则
}

// 获取已设置的二维码规则
func GetQrcodeJump(clt *core.Client)(isopen, quota, count int, list []*JumpRule, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumpget?access_token="

	var result struct {
		core.Error
		QrcodejumpOpen int `json:"qrcodejump_open"`
		QrcodejumpPubQuota int `json:"qrcodejump_pub_quota"`
		ListSize int `json:"list_size"`
		RuleList []*JumpRule `json:"rule_list"`
	}

	if err = clt.PostJSON(incompleteURL, &map[string]string{}, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	isopen = result.QrcodejumpOpen
	quota = result.QrcodejumpPubQuota
	count = result.ListSize
	list = result.RuleList
	return
}

// 获取校验文件名称及内容
func DownloadQrcodeJump(clt *core.Client)(filename, filecontent string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumpdownload?access_token="

	var result struct {
		core.Error
		FileName string `json:"file_name"`
		FileContent string `json:"file_content"`
	}

	if err = clt.PostJSON(incompleteURL, &map[string]string{}, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	filename = result.FileName
	filecontent = result.FileContent
	return
}

// 增加或修改二维码规则
func AddQrcodeJump(clt *core.Client, rule *JumpRule)(err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumpadd?access_token="

	var result struct {
		core.Error
	}

	if err = clt.PostJSON(incompleteURL, rule, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}

// 发布已设置的二维码规则
func PublishQrcodeJump(clt *core.Client, prefix string)(err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumppublish?access_token="

	var result struct {
		core.Error
	}

	if err = clt.PostJSON(incompleteURL, &map[string]string{"prefix": prefix}, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}

// 发布已设置的二维码规则
func DeleteQrcodeJump(clt *core.Client, prefix string)(filename, filecontent string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumpdelete?access_token="

	var result struct {
		core.Error
	}

	if err = clt.PostJSON(incompleteURL, &map[string]string{"prefix": prefix}, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}