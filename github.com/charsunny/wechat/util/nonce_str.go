package util

import "github.com/charsunny/wechat/rand"

func NonceStr() string {
	return string(rand.NewHex())
}
