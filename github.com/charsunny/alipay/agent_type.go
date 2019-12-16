package alipay

type AgentCreateContactInfo struct {
	ContactName     string `json:"contact_name"` // 代商家创建的小程序应用名称。名称可以由中文、数字、英文及下划线组成，长度在3-20个字符之间
	ContactMobile     string `json:"contact_mobile"` // 小程序英文名称，长度3~20个字符
	ContactEmail     string `json:"contact_email"` // 小程序应用类目，参数格式：一级类目_二级类目 https://docs.alipay.com/isv/10325
}

// https://docs.open.alipay.com/api_50/alipay.open.agent.create/
type AgentCreateQuery struct {
	AppAuthToken string `json:"-"` // 可选
	Account     string `json:"account"` // isv代操作的商户账号，可以是支付宝账号，也可以是pid（2088开头）
	ContactInfo *AgentCreateContactInfo `json:"contact_info"`
	OrderTicket     string `json:"order_ticket"`  // 代商家创建的小程序的简介，请用一句话简要描述小程序提供的服务；应用上架后一个自然月可修改5次（10~32个字符）
}

func (this AgentCreateQuery) APIName() string {
	return "alipay.open.agent.create"
}

func (this AgentCreateQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this AgentCreateQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this AgentCreateQuery) ExtJSONParamValue() string {
	return Marshal(this)
}

type AgentCreateRsp struct {
	Content struct {
		BatchNo 		string `json:"batch_no"`
		BatchStatus 	string `json:"batch_status"`
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code,omitempty"`
		SubMsg          string `json:"sub_msg,omitempty"`
	} `json:"alipay_open_agent_create_response"`
	Sign string `json:"sign"`
}

// https://docs.open.alipay.com/api_50/alipay.open.agent.confirm/
type AgentConfirmQuery struct {
	AppAuthToken string `json:"-"` // 可选
	BatchNo 		string `json:"batch_no"` // ISV 代商户操作事务编号，通过事务开启接口alipay.open.agent.create调用返回。
}

func (this AgentConfirmQuery) APIName() string {
	return "alipay.open.agent.create"
}

func (this AgentConfirmQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this AgentConfirmQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this AgentConfirmQuery) ExtJSONParamValue() string {
	return Marshal(this)
}

type AgentConfirmRsp struct {
	Content struct {
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code,omitempty"`
		SubMsg          string `json:"sub_msg,omitempty"`
		UserId 			string `json:"user_id"`
		AuthAppId 		string `json:"auth_app_id"`
		AppAuthToken 	string `json:"app_auth_token"`
		AppRefreshToken string `json:"app_refresh_token"`
		ExpiresIn 		string `json:"expires_in"`
		ReExpiresIn 	string `json:"re_expires_in"`
	} `json:"alipay_open_agent_confirm_response"`
	Sign string `json:"sign"`
}

// https://docs.open.alipay.com/api_50/alipay.open.agent.confirm/
type AgentCancelQuery struct {
	AppAuthToken string `json:"-"` // 可选
	BatchNo 		string `json:"batch_no"` // ISV 代商户操作事务编号，通过事务开启接口alipay.open.agent.create调用返回。
}

func (this AgentCancelQuery) APIName() string {
	return "alipay.open.agent.cancel"
}

func (this AgentCancelQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this AgentCancelQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this AgentCancelQuery) ExtJSONParamValue() string {
	return Marshal(this)
}

type AgentCancelRsp struct {
	Content struct {
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code,omitempty"`
		SubMsg          string `json:"sub_msg,omitempty"`
	} `json:"alipay_open_agent_cancel_response"`
	Sign string `json:"sign"`
}

// https://docs.open.alipay.com/api_50/alipay.open.agent.mini.create/
type AgentMiniCreateQuery struct {
	AppAuthToken string `json:"-"` // 可选
	BatchNo     string `json:"batch_no"` // ISV 代商家操作事务编号，通过事务开启接口alipay.open.agent.create调用返回。
	AppName     string `json:"app_name"` // 代商家创建的小程序应用名称。名称可以由中文、数字、英文及下划线组成，长度在3-20个字符之间
	AppEnglishName     string `json:"app_english_name"` // 小程序英文名称，长度3~20个字符
	AppCategoryIds     string `json:"app_category_ids"` // 小程序应用类目，参数格式：一级类目_二级类目 https://docs.alipay.com/isv/10325
	AppSlogan     string `json:"app_slogan"`  // 代商家创建的小程序的简介，请用一句话简要描述小程序提供的服务；应用上架后一个自然月可修改5次（10~32个字符）
	ServicePhone     string `json:"service_phone,omitempty"` // 商家小程序的客服电话，推荐填写
	ServiceEmail     string `json:"service_email,omitempty"` // 商家小程序的客服电话，推荐填写
	AppDesc     string `json:"app_desc,omitempty"` // 商家小程序描述信息，简要描述小程序主要功能（20-200个字）
	AppLogo string `json:"-"`
}

func (this AgentMiniCreateQuery) APIName() string {
	return "alipay.open.agent.mini.create"
}

func (this AgentMiniCreateQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["batch_no"] = this.BatchNo
	m["app_name"] = this.AppName
	m["app_english_name"] = this.AppEnglishName
	m["app_category_ids"] = this.AppCategoryIds
	m["app_slogan"] = this.AppSlogan
	m["service_phone"] = this.ServicePhone
	m["service_email"] = this.ServiceEmail
	m["app_desc"] = this.AppDesc
	return m
}

func (this AgentMiniCreateQuery) ExtJSONParamName() string {
	return ""
}

func (this AgentMiniCreateQuery) ExtJSONParamValue() string {
	return ""
}

type AgentMiniCreateRsp struct {
	Content struct {
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code,omitempty"`
		SubMsg          string `json:"sub_msg,omitempty"`
	} `json:"alipay_open_agent_mini_create_response"`
	Sign string `json:"sign"`
}

// https://docs.open.alipay.com/api_50/alipay.open.agent.mini.create/
type AgentMiniModifyQuery struct {
	AppAuthToken string `json:"-"` // 可选
	BatchNo     string `json:"batch_no"` // ISV 代商家操作事务编号，通过事务开启接口alipay.open.agent.create调用返回。
	AppName     string `json:"app_name"` // 代商家创建的小程序应用名称。名称可以由中文、数字、英文及下划线组成，长度在3-20个字符之间
	AppEnglishName     string `json:"app_english_name"` // 小程序英文名称，长度3~20个字符
	AppCategoryIds     string `json:"app_category_ids"` // 小程序应用类目，参数格式：一级类目_二级类目 https://docs.alipay.com/isv/10325
	AppSlogan     string `json:"app_slogan"`  // 代商家创建的小程序的简介，请用一句话简要描述小程序提供的服务；应用上架后一个自然月可修改5次（10~32个字符）
	ServicePhone     string `json:"service_phone,omitempty"` // 商家小程序的客服电话，推荐填写
	ServiceEmail     string `json:"service_email,omitempty"` // 商家小程序的客服电话，推荐填写
	AppDesc     string `json:"app_desc,omitempty"` // 商家小程序描述信息，简要描述小程序主要功能（20-200个字）
	AppLogo string `json:"-"`
}

func (this AgentMiniModifyQuery) APIName() string {
	return "alipay.open.mini.baseinfo.modify"
}

func (this AgentMiniModifyQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["batch_no"] = this.BatchNo
	m["app_name"] = this.AppName
	m["app_english_name"] = this.AppEnglishName
	m["app_category_ids"] = this.AppCategoryIds
	m["app_slogan"] = this.AppSlogan
	m["service_phone"] = this.ServicePhone
	m["service_email"] = this.ServiceEmail
	m["app_desc"] = this.AppDesc
	return m
}

func (this AgentMiniModifyQuery) ExtJSONParamName() string {
	return ""
}

func (this AgentMiniModifyQuery) ExtJSONParamValue() string {
	return ""
}

type AgentMiniModifyRsp struct {
	Content struct {
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code,omitempty"`
		SubMsg          string `json:"sub_msg,omitempty"`
	} `json:"alipay_open_agent_mini_create_response"`
	Sign string `json:"sign"`
}


