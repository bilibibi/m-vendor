package alipay

// https://docs.open.alipay.com/api_50/alipay.open.agent.confirm/
type OpenAuthTokenQuery struct {
	GrantType string `json:"grant_type"` // 如果使用app_auth_code换取token，则为authorization_code，如果使用refresh_token换取新的token，则为refresh_token
	Code 		string `json:"code"` // 与refresh_token二选一，用户对应用授权后得到，即第一步中开发者获取到的app_auth_code值
	RefreshToken 		string `json:"refresh_token"` // 与code二选一，可为空，刷新令牌时使用
}

func (this OpenAuthTokenQuery) APIName() string {
	return "alipay.open.auth.token.app"
}

func (this OpenAuthTokenQuery) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

func (this OpenAuthTokenQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this OpenAuthTokenQuery) ExtJSONParamValue() string {
	if this.Code == "" {
		this.GrantType = "refresh_token"
	} else {
		this.GrantType = "authorization_code"
	}
	return Marshal(this)
}

type OpenAuthTokenRsp struct {
	Content struct {
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code,omitempty"`
		SubMsg          string `json:"sub_msg,omitempty"`
		Tokens []struct{
			AppAuthToken 	string `json:"app_auth_token"`
			UserId 			string `json:"user_id"`
			AuthAppId 		string `json:"auth_app_id"`
			ExpiresIn 		int `json:"expires_in"`
			ReExpiresIn 	int `json:"re_expires_in"`
			AppRefreshToken string `json:"app_refresh_token"`
		} `json:"tokens"`
	} `json:"alipay_open_auth_token_app_response"`
	Sign string `json:"sign"`
}

// https://docs.open.alipay.com/api_50/alipay.open.agent.confirm/
type OpenAuthStatusQuery struct {
	AppAuthToken string `json:"app_auth_token"` // 如
}

func (this OpenAuthStatusQuery) APIName() string {
	return "alipay.open.auth.token.app.query"
}

func (this OpenAuthStatusQuery) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

func (this OpenAuthStatusQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this OpenAuthStatusQuery) ExtJSONParamValue() string {
	return Marshal(this)
}

type OpenAuthStatusRsp struct {
	Content struct {
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code,omitempty"`
		SubMsg          string `json:"sub_msg,omitempty"`
		UserId 			string `json:"user_id"`
		AuthAppId 		string `json:"auth_app_id"`
		AuthMethods 	string `json:"auth_methods"`
		AuthStart 	 	string `json:"auth_start"`
		AuthEnd 		string `json:"auth_end"`
		Status 			string `json:"status"`
	} `json:"alipay_open_auth_token_app_query_response"`
	Sign string `json:"sign"`
}


// https://docs.open.alipay.com/api_50/alipay.open.agent.confirm/
type UserAuthTokenQuery struct {
	AppAuthToken string `json:"app_auth_token"` // 如
	GrantType string `json:"grant_type"` // 如果使用app_auth_code换取token，则为authorization_code，如果使用refresh_token换取新的token，则为refresh_token
	Code 		string `json:"code"` // 与refresh_token二选一，用户对应用授权后得到，即第一步中开发者获取到的app_auth_code值
	RefreshToken 		string `json:"refresh_token"` // 与code二选一，可为空，刷新令牌时使用
}

func (this UserAuthTokenQuery) APIName() string {
	return "alipay.system.oauth.token"
}

func (this UserAuthTokenQuery) Params() map[string]string {
	var m = make(map[string]string)
	if this.AppAuthToken != "" {
		m["app_auth_token"] = this.AppAuthToken
	}
	if this.Code == "" {
		m["refresh_token"] = this.RefreshToken
		m["grant_type"] = "refresh_token"
	} else {
		m["code"] = this.Code
		m["grant_type"] = "authorization_code"
	}
	return m
}

func (this UserAuthTokenQuery) ExtJSONParamName() string {
	return ""
}

func (this UserAuthTokenQuery) ExtJSONParamValue() string {
	return ""
}

type UserAuthTokenRsp struct {
	Content* struct {
		UserId 			string `json:"user_id"`
		AlipayUserId 	string `json:"alipay_user_id"`
		AccessToken 	string `json:"access_token"`
		RefreshToken 	string `json:"refresh_token"`
		ExpiresIn 		int `json:"expires_in"`
		ReExpiresIn 	int `json:"re_expires_in"`
	} `json:"alipay_system_oauth_token_response,omitempty"`
	ErrorResponse* struct{
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code,omitempty"`
		SubMsg          string `json:"sub_msg,omitempty"`
	} `json:"error_response,omitempty"`
	Sign string `json:"sign"`
}

// https://docs.open.alipay.com/api_50/alipay.open.agent.confirm/
type UserAuthInfoQuery struct {
	AppAuthToken string `json:"app_auth_token"` // 如
	AuthToken 		string `json:"auth_token"` // 与code二选一，可为空，刷新令牌时使用
}

func (this UserAuthInfoQuery) APIName() string {
	return "alipay.user.info.share"
}

func (this UserAuthInfoQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["auth_token"] = this.AuthToken
	return m
}

func (this UserAuthInfoQuery) ExtJSONParamName() string {
	return ""
}

func (this UserAuthInfoQuery) ExtJSONParamValue() string {
	return ""
}

type UserAuthInfoRsp struct {
	Content struct {
		Code            	string `json:"code"`
		Msg             	string `json:"msg"`
		SubCode         	string `json:"sub_code,omitempty"`
		SubMsg          	string `json:"sub_msg,omitempty"`
		UserId 				string `json:"user_id"`
		Avatar 				string `json:"avatar"`
		Province 			string `json:"province"`
		City 				string `json:"nick_name"`
		NickName 			string `json:"nick_name"`
		Gender 				string `json:"gender"`		// 只有is_certified为T的时候才有意义，否则不保证准确性. 性别（F：女性；M：男性）
		IsStudentCertified 	string `json:"is_student_certified"`
		IsCertified 		string `json:"is_certified"`// 是否通过实名认证。T是通过 F是没有实名认证。
		UserType 			string `json:"user_type"`	// 用户类型（1/2） 1代表公司账户2代表个人账户
		UserStatus			string `json:"user_status"`	//用户状态（Q/T/B/W）。 Q代表快速注册用户 T代表已认证用户 B代表被冻结账户 W代表已注册，未激活的账户
	} `json:"alipay_user_info_share_response"`
	Sign string `json:"sign"`
}
