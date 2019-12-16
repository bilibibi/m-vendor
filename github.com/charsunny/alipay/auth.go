package alipay

import (
	"fmt"
	"net/url"
)

func (this *Client) GetOpenAuthURL(redirect_uri string) string {
	return fmt.Sprintf("https://openauth.alipay.com/oauth2/appToAppAuth.htm?app_id=%s&redirect_uri=%s", this.appId, url.QueryEscape(redirect_uri))
}

func (this *Client) GetOpenAuthToken(param OpenAuthTokenQuery) (results *OpenAuthTokenRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

func (this *Client) QueryOpenAuthStatus(param OpenAuthStatusQuery) (results *OpenAuthStatusRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

func (this *Client) GetUserAuthToken(param UserAuthTokenQuery) (results *UserAuthTokenRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

func (this *Client) QueryAuthUserInfo(param UserAuthInfoQuery) (results *UserAuthInfoRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}