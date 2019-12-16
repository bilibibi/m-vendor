package merchant

// DatabaseTable

//  商品
type Product struct {
	ProductId string `json:"product_id"`
	ProductBase *ProductBase `json:"product_base"`
	SkuList []*SkuItem	`json:"sku_list"`
	AttrExt *AttrExt 	`json:"attrext"`
	DeliveryInfo *DeliveryInfo `json:"delivery_info"`
}

type ProductBase struct {
	CategoryId []int                     `json:"category_id"`
	Property   []* struct {
		Id  string `json:"id"`
		Vid string `json:"vid"`
	}          `json:"property"`
	Name       string                       `json:"name"`
	SkuInfo    []*SkuInfo           `json:"sku_info"`
	MainImg    string                       `json:"main_img"`
	Img        []string                     `json:"img"`
	Detail     []*ProductDetail `json:"detail"`
	DetailHtml string `json:"detail_html"`
	BuyLimit   int                          `json:"buy_limit"`
}

// SkuItem 库存
type SkuItem struct {
	SkuId       string `json:"sku_id"`
	Price       int64  `json:"price"`
	IconUrl     string `json:"icon_url"`
	ProductCode string `json:"product_code"`
	OriPrice    int64  `json:"ori_price"`
	Quantity    int    `json:"quantity"`
}

// Express 快递模板
type Express struct {
	TemplateId int    `json:"Id"`
	Name      string       `json:"Name"`		// 邮费模板名称
	Assumer   int          `json:"Assumer"`		// 支付方式(0-买家承担运费, 1-卖家承担运费)
	Valuation int          `json:"Valuation"`	// 计费单位(0-按件计费, 1-按重量计费, 2-按体积计费，目前只支持按件计费，默认为0)
	TopFee    []*TopFeeItem `json:"TopFee"`	// 具体运费计算
}

// Express 快递费用模板
type TopFeeItem struct {
	Type   int `json:"Type"`		// 快递类型ID(参见增加商品/快递列表)
	Normal struct {	// 默认邮费计算方法
		StartStandards int `json:"StartStandards"`	// 起始计费数量(比如计费单位是按件, 填2代表起始计费为2件)
		StartFees      int `json:"StartFees"`	// 起始计费金额(单位: 分）
		AddStandards   int `json:"AddStandards"`	// 递增计费数量
		AddFees        int `json:"AddFees"`	// 递增计费金额(单位 : 分)
	} `json:"Normal"`
	Custom []struct {	// 指定地区邮费计算方法
		StartStandards int    `json:"StartStandards"`
		StartFees      int    `json:"StartFees"`
		AddStandards   int    `json:"AddStandards"`
		AddFees        int    `json:"AddFees"`
		DestCountry    string `json:"DestCountry"`	// 指定国家(详见《地区列表》说明)
		DestProvince   string `json:"DestProvince"`	// 指定省份(详见《地区列表》说明)
		DestCity       string `json:"DestCity"`	// 指定城市(详见《地区列表》说明)
	} `json:"Custom"`
}

type SkuInfo struct {
	Id  string   `json:"id"`
	Vid []string `json:"vid"`
}

type ProductDetail struct {
	Text string `json:"text"`
	Img  string `json:"img`
}

type SkuListItem struct {
	SkuId       string `json:"sku_id"`
	Price       int64  `json:"price"`
	IconUrl     string `json:"icon_url"`
	ProductCode string `json:"product_code"`
	OriPrice    int64  `json:"ori_price"`
	Quantity    int    `json:"quantity"`
}

type AttrExtLocation struct {
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	Address  string `json:"address"`
}

type AttrExt struct {
	Location         *AttrExtLocation `json:"location"`
	IsPostFree       int                      `json:"isPostFree"`
	IsHasReceipt     int                      `json:"isHasReceipt"`
	IsUnderGuaranty  int                      `json:"isUnderGuaranty"`
	IsSupportReplace int                      `json:"isSupportReplace"`
}

type DeliveryInfo struct {
	DeliveryType int                            `json:"delivery_type"`	// 运费类型(0-使用下面express字段的默认模板, 1-使用template_id代表的邮费模板, 详见邮费模板相关API)
	TemplateId   int                            `json:"template_id"`
	Express      []*DeliveryInfoExpress `json:"express"`
}

type DeliveryInfoExpress struct {
	Id    int64 `json:"id"`
	Price int64 `json:"price"`
}


type CateInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SkuValueList struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SkuTable struct {
	Id        string                  `json:"id"`
	Name      string                  `json:"name"`
	ValueList []SkuValueList `json:"value_list"`
}

type CateProperty struct {
	Id            string                           `json:"id"`
	Name          string                           `json:"name"`
	PropertyValue []*CateProperty `json:"property_value"`
}

type PropertyValue struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type StockSkuInfo struct {
	Id  string
	Vid string
}

type GroupInfo struct {
	GroupId   string `json:"group_id"`
	GroupName string `json:"group_name"`
	ProductList []string `json:"product_list"`
}
