package wxa

type WxaInfo struct {
	Appid string `json:"appid"`	// 小程序gh_id
	AccountType int `json:"account_type"`//帐号类型（1：订阅号，2：服务号，3：小程序）；
	PrincipalType int `json:"principal_type"` // 主体类型（1：企业）
	PrincipalName string `json:"principal_name"`	// 主体名称
	RealNameStatus int `json:"realname_status"`	// 实名验证状态（1：实名验证成功，2：实名验证中，3：实名验证失败）调用接口1.1创建帐号时，realname_status会初始化为2对于注册方式为微信认证的帐号，资质认证成功时，realname_status会更新为1 注意！！！当realname_status不为1时，帐号只允许调用本文档内的以下API：（即无权限调用其他API） 微信认证相关接口（参考2.x） 帐号设置相关接口（
	WxVerifyInfo struct {
		QualificationVerify bool `json:"qualification_verify"`	//是否资质认证（true：是，false：否）若是，拥有微信认证相关的权限
		NamingVerify bool `json:"naming_verify"` // 是否名称认证（true：是，false：否）对于公众号（订阅号、服务号），是名称认证，微信客户端才会有微信认证打勾的标识
		AnnualReview bool `json:"annual_review"`	// 是否需要年审（true：是，false：否）（qualification_verify = true时才有该字段）
		AnnualReviewBeginTime int `json:"annual_review_begin_time"`	// 年审开始时间，时间戳（qualification_verify = true时才有该字段）
		AnnualReviewEndTime int `json:"annual_review_end_time"` // 年审截止时间，时间戳（qualification_verify = true时才有该字段）
	} `json:"wx_verify_info"`
	SignatureInfo struct {
		Signature string `json:"signature"`	// 功能介绍
		ModifyUsedCount int `json:"modify_used_count"`	// 功能介绍已使用修改次数（本月）
		ModifyQuota int `json:"modify_quota"`	// 功能介绍修改次数总额度（本月）
	} `json:"signature_info"`
	HeadImageInfo struct{
		HeadImageUrl string `json:"head_image_url"`
		ModifyUsedCount int `json:"modify_used_count"`	// 功能介绍已使用修改次数（本月）
		ModifyQuota int `json:"modify_quota"`	// 功能介绍修改次数总额度（本月）
	} `json:"head_image_info"`
}

type WxaActionCateInfo struct {
	First int `json:"first"`
	Second int `json:"second"`
	Certicates [] struct {
		Key string `json:"key"`
		Value string `json:"value"`
	} `json:"certicates"`
}

type WxaCategory struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Father int `json:"father"`
	Level int `json:"level"`
	Children []int `json:"children"`
	SensitiveType int `json:"sensitive_type"`
	Qualify struct {
		ExterList [] struct{
			InnerList [] struct{
				Name string `json:"name"`
				Url string `json:"url"`
			} `json:"inner_list"`
		} `json:"exter_list"`
	} `json:"qualify"`
}

type CategoryItem struct {
	Quota int `json:"quota"`
	Limit int `json:"limit"`
	CategoryLimit int `json:"category_limit"`
	Categories [] struct{
		First int `json:"first"`
		FirstName string `json:"first_name"`
		Second int `json:"second"`
		SecondName string `json:"second_name"`
		AuditStatus int `json:"audit_status"`
		AuditReason string `json:"audit_reason"`
	} `json:"categories"`
}

type WxaNameRequestInfo struct {
	NickName string `json:"nick_name"`	// 昵称
	IdCard string `json:"id_card"`	// 身份证照片–临时素材mediaid	个人号必填
	License string `json:"license"`	// 组织机构代码证或营业执照–临时素材mediaid	组织号必填
	NamingOtherStuff1 string `json:"naming_other_stuff_1"`	// 其他证明材料---临时素材 mediaid	选填
	NamingOtherStuff2 string `json:"naming_other_stuff_2"`	// 其他证明材料---临时素材 mediaid	选填
}

type WxaNameResultInfo struct {
	NickName string `json:"nickname"`	// 昵称
	AuditStat int `json:"audit_stat"`	// 审核状态，1：审核中，2：审核失败，3：审核成功
	FailReason string `json:"fail_reason"`	// 失败原因
	CreateTime int `json:"create_time"`	// 审核提交时间
	AuditTime int `json:"audit_time"`	// 审核完成时间
}

type WxaItemInfo struct {
	AppId string `json:"appid"`	// 展示的公众号appid
	Nickname string `json:"nickname"`	// 展示的公众号nickname
	Headimg string `json:"headimg"`	// 展示的公众号头像
	CanOpen int `json:"can_open"`	// 是否可以设置 1 可以，0，不可以
	IsOpen int `json:"is_open"`	// 是否已经设置，1 已设置，0，未设置
}

type WxaCategoryInfo struct {
	FirstClass string `json:"first_class"`
	SecondClass string `json:"second_class"`
	ThirdClass string `json:"third_class"`
	FirstId int `json:"first_id"`
	SecondId int `json:"second_id"`
	ThirdId int `json:"third_id"`
}

type WxaPageInfo struct {
	Title string `json:"title"`
	Address string `json:"address"`
	Tag string `json:"tag"`
	WxaCategoryInfo
}
