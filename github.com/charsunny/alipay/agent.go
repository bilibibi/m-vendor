package alipay

func (this *Client)AgentCreate(param AgentCreateQuery) (results *AgentCreateRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

func (this *Client)AgentConfirm(param AgentConfirmQuery) (results *AgentConfirmRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

func (this *Client)AgentCancel(param AgentCancelQuery) (results *AgentCancelRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// BillDownloadURLQuery https://docs.open.alipay.com/api_15/alipay.data.dataservice.bill.downloadurl.query
func (this *Client)AgentMiniCreate(param AgentMiniCreateQuery) (results *AgentMiniCreateRsp, err error) {
	files := map[string]string {
		"app_logo": param.AppLogo,
	}
	err = this.doRequestWithFile("POST", param,  files, &results)
	return results, err
}

func (this *Client)AgentMiniModify(param AgentMiniCreateQuery) (results *AgentMiniModifyRsp, err error) {
	files := map[string]string {
		"app_logo": param.AppLogo,
	}
	err = this.doRequestWithFile("POST", param,  files, &results)
	return results, err
}


