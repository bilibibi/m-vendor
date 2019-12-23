package pay

import (
	"fmt"
	"github.com/charsunny/wechat/mch/core"
	"github.com/charsunny/wechat/util"
	"time"
)


type FacepayAuthInfoRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	StoreId string `xml:"store_id"` // 门店编号， 由商户定义， 各门店唯一。
	StoreName    string `xml:"store_name"`   // 门店名称，由商户定义。（可用于展示）
	DeviceId    string `xml:"device_id"`   // 终端设备编号，由商户定义。
	Rawdata string `xml:"rawdata"`   // 初始化数据。由微信人脸SDK的接口返回。

	// 可选参数
	Attach string `xml:"attach"` // 附加字段。字段格式使用Json
	NonceStr string `xml:"nonce_str"` // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
	SignType string `xml:"sign_type"` // 签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
}

// Reverse2 撤销订单.
//  NOTE: 请求需要双向证书.
func GetWxpayfaceAuthinfo(clt *core.Client, req *FacepayAuthInfoRequest) (resp map[string]string, err error) {
	m1 := make(map[string]string, 8)
	m1["store_id"] = req.StoreId
	m1["store_name"] = req.StoreName
	m1["DeviceId"] = req.DeviceId
	m1["rawdata"] = req.Rawdata
	m1["version"] = "1"
	m1["now"] = fmt.Sprintf("%010d", time.Now().Unix())
	if req.NonceStr != "" {
		m1["nonce_str"] = req.NonceStr
	} else {
		m1["nonce_str"] = util.NonceStr()
	}
	if req.SignType != "" {
		m1["sign_type"] = req.SignType
	}

	resp, err = clt.PostXML(core.APIAppURL()+"/face/get_wxpayface_authinfo", m1)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
