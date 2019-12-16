package analysis

import (
	"github.com/charsunny/wechat/mp/core"
)

// 获取小程序新增或活跃用户的画像分布数据。时间范围支持昨天、最近7天、最近30天。
// @param begin_date	string		是	开始日期。格式为 yyyymmdd
// @param end_date	string		结束日期，开始日期与结束日期相差的天数限定为0/6/29，分别表示查询最近1/7/30天数据，
// @return visitUvNew	Object	新增用户留存
// @return visitUv	Object	新增用户留存
func GetUserPortrait(clt *core.Client, begin_date string, end_date string) (date string, visit_new map[string]interface{}, visit map[string]interface{}, err error) {
	const incompleteURL = "https://api.weixin.qq.com/datacube/getweanalysisappiduserportrait?access_token="

	var req = struct {
		BeginDate string `json:"begin_date"`
		EndDate string `json:"end_date"`
	} {
		BeginDate:begin_date,
		EndDate:end_date,
	}
	var result struct {
		core.Error
		RefDate string `json:"ref_date"`
		VisitUvNew map[string]interface{} `json:"visit_uv_new"`
		VisitUv map[string]interface{} `json:"visit_uv"`
	}
	if err = clt.PostJSON(incompleteURL, &req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	date = result.RefDate
	visit_new = result.VisitUvNew
	visit = result.VisitUv
	return
}

// 获取用户小程序访问分布数据
// @param begin_date	string		是	开始日期。格式为 yyyymmdd
// @param end_date	string		结束日期，限定查询 1 天数据，允许设置的最大值为昨日。格式为 yyyymmdd，
func GetVisitDistribution(clt *core.Client, begin_date string, end_date string)(date string, list []map[string]interface{}, err error) {
	const incompleteURL = "https://api.weixin.qq.com/datacube/getweanalysisappidvisitdistribution?access_token="

	var req = struct {
		BeginDate string `json:"begin_date"`
		EndDate string `json:"end_date"`
	} {
		BeginDate:begin_date,
		EndDate:end_date,
	}
	var result struct {
		core.Error
		RefDate string `json:"ref_date"`
		List []map[string]interface{} `json:"list"`
	}
	if err = clt.PostJSON(incompleteURL, &req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	date = result.RefDate
	list = result.List
	return
}

// 获取用户小程序访问分布数据
// @param begin_date	string		是	开始日期。格式为 yyyymmdd
// @param end_date	string		结束日期，限定查询 1 天数据，允许设置的最大值为昨日。格式为 yyyymmdd，
func GetVisitPage(clt *core.Client, begin_date string, end_date string)(date string, list []map[string]interface{}, err error) {
	const incompleteURL = "https://api.weixin.qq.com/datacube/getweanalysisappidvisitpage?access_token="

	var req = struct {
		BeginDate string `json:"begin_date"`
		EndDate string `json:"end_date"`
	} {
		BeginDate:begin_date,
		EndDate:end_date,
	}
	var result struct {
		core.Error
		RefDate string `json:"ref_date"`
		List []map[string]interface{} `json:"list"`
	}
	if err = clt.PostJSON(incompleteURL, &req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	date = result.RefDate
	list = result.List
	return
}

// 获取用户小程序访问分布数据
// @param begin_date	string		是	开始日期。格式为 yyyymmdd
// @param end_date	string		结束日期，限定查询 1 天数据，允许设置的最大值为昨日。格式为 yyyymmdd，
func GetDailySummary(clt *core.Client, begin_date string, end_date string)(list []map[string]interface{}, err error) {
	const incompleteURL = "https://api.weixin.qq.com/datacube/getweanalysisappiddailysummarytrend?access_token="

	var req = struct {
		BeginDate string `json:"begin_date"`
		EndDate string `json:"end_date"`
	} {
		BeginDate:begin_date,
		EndDate:end_date,
	}
	var result struct {
		core.Error
		List []map[string]interface{} `json:"list"`
	}
	if err = clt.PostJSON(incompleteURL, &req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.List
	return
}