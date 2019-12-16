package alipay

// https://docs.open.alipay.com/api_49/alipay.open.mini.version.upload
type MiniVersionUploadQuery struct {
	AppAuthToken 	string `json:"-"` // 可选
	TemplateVersion string `json:"template_version"` // 可选, 模板版本号，版本号必须满足 x.y.z, 且均为数字
	Ext 			string `json:"ext"`				// 模板的配置参数	{"extEnable": true, "extPages": {"pages/face/index": {"defaultTitle": "哈哈哈哈"}},"window": {"defaultTitle": "AI2"}}
	TemplateId     	string `json:"template_id"`  	// 模板id	1
	AppVersion		string `json:"app_version"`		// 小程序版本号，版本号必须满足 x.y.z, 且均为数字
}

func (this MiniVersionUploadQuery) APIName() string {
	return "alipay.open.mini.version.upload"
}

func (this MiniVersionUploadQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this MiniVersionUploadQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this MiniVersionUploadQuery) ExtJSONParamValue() string {
	return Marshal(this)
}

type MiniVersionUploadRsp struct {
	Content struct {
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code,omitempty"`
		SubMsg          string `json:"sub_msg,omitempty"`
	} `json:"alipay_open_mini_version_upload_response"`
	Sign string `json:"sign"`
}

// https://docs.open.alipay.com/api_49/alipay.open.mini.version.build
type MiniVersionBuildQuery struct {
	AppAuthToken 	string `json:"-"` // 可选
	AppVersion		string `json:"app_version"`		// 小程序版本号，版本号必须满足 x.y.z, 且均为数字
}

func (this MiniVersionBuildQuery) APIName() string {
	return "alipay.open.mini.version.build.query"
}

func (this MiniVersionBuildQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this MiniVersionBuildQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this MiniVersionBuildQuery) ExtJSONParamValue() string {
	return Marshal(this)
}

type MiniVersionBuildRsp struct {
	Content struct {
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code,omitempty"`
		SubMsg          string `json:"sub_msg,omitempty"`
		NeedRotation 	string 	`json:"need_rotation"`	// 是否需要轮询
		CreateStatus 	string 	`json:"create_status"`	// 0-构建排队中；1-正在构建；2-构建成功；3-构建失败；5-构建超时；6-版本创建成功
	} `json:"alipay_open_mini_version_build_query_response"`
	Sign string `json:"sign"`
}

// https://docs.open.alipay.com/api_49/alipay.open.mini.version.delete
type MiniVersionDeleteQuery struct {
	AppAuthToken 	string `json:"-"` // 可选
	AppVersion		string `json:"app_version"`		// 小程序版本号，版本号必须满足 x.y.z, 且均为数字
}

func (this MiniVersionDeleteQuery) APIName() string {
	return "alipay.open.mini.version.delete"
}

func (this MiniVersionDeleteQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this MiniVersionDeleteQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this MiniVersionDeleteQuery) ExtJSONParamValue() string {
	return Marshal(this)
}

type MiniVersionDeleteRsp struct {
	Content struct {
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code,omitempty"`
		SubMsg          string `json:"sub_msg,omitempty"`
	} `json:"alipay_open_mini_version_delete_response"`
	Sign string `json:"sign"`
}

// https://docs.open.alipay.com/api_49/alipay.open.mini.version.audit.apply
type MiniVersionAuditApplyQuery struct {
	AppAuthToken 		string `json:"-"` // 可选
	AppVersion			string `json:"app_version"`		// 小程序版本号，版本号必须满足 x.y.z, 且均为数字
	VersionDesc			string `json:"version_desc"`		// 小程序版本号，版本号必须满足 x.y.z, 且均为数字
	LicenseNo			string `json:"license_no"`		// 营业执照号，部分小程序类目需要提交，参照https://docs.alipay.com/isv/10325
	LicenseName			string `json:"license_name"`		// 营业执照名称，部分小程序类目需要提交，参照https://docs.alipay.com/isv/10325
	LicenseValidDate 	string `json:"license_valid_date"`		// 营业执照有效期，格式为yyyy-MM-dd，9999-12-31表示长期
	FirstLicensePic		string `json:"-"`				// 第一张营业执照照片，不能超过4MB
	OutDoorPic 			string `json:"-"` // 门头照图片，部分小程序类目需要提交
}

func (this MiniVersionAuditApplyQuery) APIName() string {
	return "alipay.open.mini.version.audit.apply"
}

func (this MiniVersionAuditApplyQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["version_desc"] = this.VersionDesc
	if this.LicenseNo != "" {
		m["license_no"] = this.LicenseNo
		m["license_name"] = this.LicenseName
		m["license_valid_date"] = this.LicenseValidDate
	}
	return m
}

func (this MiniVersionAuditApplyQuery) ExtJSONParamName() string {
	return ""
}

func (this MiniVersionAuditApplyQuery) ExtJSONParamValue() string {
	return ""
}

type MiniVersionAuditApplyRsp struct {
	Content struct {
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code,omitempty"`
		SubMsg          string `json:"sub_msg,omitempty"`
	} `json:"alipay_open_mini_version_audit_apply_response"`
	Sign string `json:"sign"`
}

// https://docs.open.alipay.com/api_49/alipay.open.mini.version.detail.query
type MiniVersionDetailQuery struct {
	AppAuthToken 		string `json:"-"` // 可选
	AppVersion			string `json:"app_version"`		// 小程序版本号，版本号必须满足 x.y.z, 且均为数字
}

func (this MiniVersionDetailQuery) APIName() string {
	return "alipay.open.mini.version.detail.query"
}

func (this MiniVersionDetailQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this MiniVersionDetailQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this MiniVersionDetailQuery) ExtJSONParamValue() string {
	return Marshal(this)
}

type MiniVersionDetailRsp struct {
	Content struct {
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code,omitempty"`
		SubMsg          string `json:"sub_msg,omitempty"`
		AppName     string `json:"app_name"` // 代商家创建的小程序应用名称。名称可以由中文、数字、英文及下划线组成，长度在3-20个字符之间
		AppEnglishName     string `json:"app_english_name"` // 小程序英文名称，长度3~20个字符
		AppCategoryIds     string `json:"app_category_ids"` // 小程序应用类目，参数格式：一级类目_二级类目 https://docs.alipay.com/isv/10325
		AppSlogan     string `json:"app_slogan"`  // 代商家创建的小程序的简介，请用一句话简要描述小程序提供的服务；应用上架后一个自然月可修改5次（10~32个字符）
		ServicePhone     string `json:"service_phone,omitempty"` // 商家小程序的客服电话，推荐填写
		ServiceEmail     string `json:"service_email,omitempty"` // 商家小程序的客服电话，推荐填写
		AppDesc     string `json:"app_desc,omitempty"` // 商家小程序描述信息，简要描述小程序主要功能（20-200个字）
		Status string `json:"status"`
	} `json:"alipay_open_mini_version_detail_query_response"`
	Sign string `json:"sign"`
}

// https://docs.open.alipay.com/api_49/alipay.open.mini.version.detail.query
type MiniVersionAuditCancelQuery struct {
	AppAuthToken 		string `json:"-"` // 可选
	AppVersion			string `json:"app_version"`		// 小程序版本号，版本号必须满足 x.y.z, 且均为数字
}

func (this MiniVersionAuditCancelQuery) APIName() string {
	return "alipay.open.mini.version.audited.cancel"
}

func (this MiniVersionAuditCancelQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this MiniVersionAuditCancelQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this MiniVersionAuditCancelQuery) ExtJSONParamValue() string {
	return Marshal(this)
}

type MiniVersionAuditCancelRsp struct {
	Content struct {
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code,omitempty"`
		SubMsg          string `json:"sub_msg,omitempty"`
	} `json:"alipay_open_mini_version_audited_cancel_response"`
	Sign string `json:"sign"`
}

// https://docs.open.alipay.com/api_49/alipay.open.mini.experience.create
type MiniExperienceCreateQuery struct {
	AppAuthToken 		string `json:"-"` // 可选
	AppVersion			string `json:"app_version"`		// 小程序版本号，版本号必须满足 x.y.z, 且均为数字
}

func (this MiniExperienceCreateQuery) APIName() string {
	return "alipay.open.mini.experience.create"
}

func (this MiniExperienceCreateQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this MiniExperienceCreateQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this MiniExperienceCreateQuery) ExtJSONParamValue() string {
	return Marshal(this)
}

type MiniExperienceCreateRsp struct {
	Content struct {
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code,omitempty"`
		SubMsg          string `json:"sub_msg,omitempty"`
	} `json:"alipay_open_mini_experience_create_response"`
	Sign string `json:"sign"`
}

// https://docs.open.alipay.com/api_49/alipay.open.mini.experience.create
type MiniExperienceQrcodeQuery struct {
	AppAuthToken 		string `json:"-"` // 可选
	AppVersion			string `json:"app_version"`		// 小程序版本号，版本号必须满足 x.y.z, 且均为数字
}

func (this MiniExperienceQrcodeQuery) APIName() string {
	return "alipay.open.mini.experience.query"
}

func (this MiniExperienceQrcodeQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this MiniExperienceQrcodeQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this MiniExperienceQrcodeQuery) ExtJSONParamValue() string {
	return Marshal(this)
}

type MiniExperienceQrcodeRsp struct {
	Content struct {
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code,omitempty"`
		SubMsg          string `json:"sub_msg,omitempty"`
		QrcodeUrl string `json:"exp_qr_code_url"`
	} `json:"alipay_open_mini_experience_query_response"`
	Sign string `json:"sign"`
}

// https://docs.open.alipay.com/api_5/alipay.open.app.qrcode.create
type MiniQrcodeCreateQuery struct {
	AppAuthToken 		string `json:"-"` // 可选
	UrlParam			string `json:"url_param"`		// 小程序版本号，版本号必须满足 x.y.z, 且均为数字
	QueryParam			string `json:"query_param"`		// 小程序版本号，版本号必须满足 x.y.z, 且均为数字
	Describe			string `json:"describe"`		// 小程序版本号，版本号必须满足 x.y.z, 且均为数字
}

func (this MiniQrcodeCreateQuery) APIName() string {
	return "alipay.open.app.qrcode.create"
}

func (this MiniQrcodeCreateQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this MiniQrcodeCreateQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this MiniQrcodeCreateQuery) ExtJSONParamValue() string {
	return Marshal(this)
}

type MiniQrcodeCreateRsp struct {
	Content struct {
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code,omitempty"`
		SubMsg          string `json:"sub_msg,omitempty"`
		QrcodeUrl string `json:"qr_code_url"`
	} `json:"alipay_open_mini_experience_create_response"`
	Sign string `json:"sign"`
}

// https://docs.open.alipay.com/api_49/alipay.open.mini.version.detail.query
type MiniVersionOnlineQuery struct {
	AppAuthToken 		string `json:"-"` // 可选
	AppVersion			string `json:"app_version"`		// 小程序版本号，版本号必须满足 x.y.z, 且均为数字
}

func (this MiniVersionOnlineQuery) APIName() string {
	return "alipay.open.mini.version.online"
}

func (this MiniVersionOnlineQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this MiniVersionOnlineQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this MiniVersionOnlineQuery) ExtJSONParamValue() string {
	return Marshal(this)
}

type MiniVersionOnlineRsp struct {
	Content struct {
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code,omitempty"`
		SubMsg          string `json:"sub_msg,omitempty"`
	} `json:"alipay_open_mini_version_online_response"`
	Sign string `json:"sign"`
}

// https://docs.open.alipay.com/api_49/alipay.open.mini.version.detail.query
type MiniVersionOfflineQuery struct {
	AppAuthToken 		string `json:"-"` // 可选
	AppVersion			string `json:"app_version"`		// 小程序版本号，版本号必须满足 x.y.z, 且均为数字
}

func (this MiniVersionOfflineQuery) APIName() string {
	return "alipay.open.mini.version.offline"
}

func (this MiniVersionOfflineQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this MiniVersionOfflineQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this MiniVersionOfflineQuery) ExtJSONParamValue() string {
	return Marshal(this)
}

type MiniVersionOfflineRsp struct {
	Content struct {
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code,omitempty"`
		SubMsg          string `json:"sub_msg,omitempty"`
	} `json:"alipay_open_mini_version_offline_response"`
	Sign string `json:"sign"`
}

// https://docs.open.alipay.com/api_49/alipay.open.mini.category.query
type MiniCategoryQuery struct {
	AppAuthToken 		string `json:"-"` // 可选
}

func (this MiniCategoryQuery) APIName() string {
	return "alipay.open.mini.category.query"
}

func (this MiniCategoryQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this MiniCategoryQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this MiniCategoryQuery) ExtJSONParamValue() string {
	return Marshal(this)
}

type CategoryList struct {
	CategoryId string `json:"category_id"`
	CategoryName string `json:"category_name"`
	ParentCategoryId string `json:"parent_category_id"`
	HasChild bool `json:"has_child"`
	NeedLicense bool `json:"need_license"`
	NeedOutDoorPic bool `json:"need_out_door_pic"`
	NeedSpecialLicense string `json:"need_special_license"`
}

type MiniCategoryRsp struct {
	Content struct {
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code,omitempty"`
		SubMsg          string `json:"sub_msg,omitempty"`
		CategoryList   []*CategoryList `json:"category_list"`
	} `json:"alipay_open_mini_category_query_response"`
	Sign string `json:"sign"`
}