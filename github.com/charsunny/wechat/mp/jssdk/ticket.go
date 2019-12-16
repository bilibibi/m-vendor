package jssdk

import (
	"github.com/charsunny/wechat/mp/core"
)

func GetSDKTickect(client *core.Client) (ticket string, expiresIn int64, err error) {
	var incompleteURL = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=jsapi&access_token="
	var result struct {
		core.Error
		Ticket    string `json:"ticket"`
		ExpiresIn int64  `json:"expires_in"`
	}
	if err = client.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	ticket = result.Ticket
	expiresIn = result.ExpiresIn
	return
}

func GetCardTickect(client *core.Client) (ticket string, expiresIn int64, err error) {
	var incompleteURL = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=wx_card&access_token="
	var result struct {
		core.Error
		Ticket    string `json:"ticket"`
		ExpiresIn int64  `json:"expires_in"`
	}
	if err = client.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	ticket = result.Ticket
	expiresIn = result.ExpiresIn
	return
}
