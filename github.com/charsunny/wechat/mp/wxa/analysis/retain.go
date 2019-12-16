package analysis

import (
	"github.com/charsunny/wechat/mp/core"
)

type VisitKV struct {
	Key int `json:"key"`
	Value int `json:"value"`
}

// 获取用户访问小程序日留存
// @param begin_date	string		是	开始日期。格式为 yyyymmdd
// @param end_date	string		是	开始日期。格式为 结束日期，限定查询1天数据
// @return visitUvNew	Object	新增用户留存
// @return visitUv	Object	新增用户留存
func GetDailyRetain(clt *core.Client, begin_date string, end_date string) (date string, visit_new []*VisitKV,  visit_uv []*VisitKV, err error) {
	const incompleteURL = "https://api.weixin.qq.com/datacube/getweanalysisappiddailyretaininfo?access_token="

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
		VisitUvNew []*VisitKV `json:"visit_uv_new"`
		VisitUv  []*VisitKV `json:"visit_uv"`
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
	visit_uv = result.VisitUv
	return
}

// 获取用户访问小程序周留存
// @param begin_date	string		是	开始日期，为周一日期。格式为 yyyymmdd
// @param end_date	string		是	开始日期。结束日期，为周日日期，限定查询一周数据。格式为 yyyymmdd
// @return visitUvNew	Object	新增用户留存
// @return visitUv	Object	新增用户留存
func GetWeeklyRetain(clt *core.Client, begin_date string, end_date string) (date string, visit_new []*VisitKV,  visit_uv []*VisitKV, err error) {
	const incompleteURL = "https://api.weixin.qq.com/datacube/getweanalysisappidweeklyretaininfo?access_token="

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
		VisitUvNew []*VisitKV `json:"visit_uv_new"`
		VisitUv  []*VisitKV `json:"visit_uv"`
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
	visit_uv = result.VisitUv
	return
}

// 获取用户访问小程序月留存
// @param begin_date	string		是	开始日期，为自然月第一天。格式为 yyyymmdd
// @param end_date	string		是	开始日期。结束日期，为自然月最后一天，限定查询一个月数据。格式为 yyyymmdd
// @return visitUvNew	Object	新增用户留存
// @return visitUv	Object	新增用户留存
func GetMonthlyRetain(clt *core.Client, begin_date string, end_date string) (date string, visit_new []*VisitKV,  visit_uv []*VisitKV, err error) {
	const incompleteURL = "https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyretaininfo?access_token="

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
		VisitUvNew []*VisitKV `json:"visit_uv_new"`
		VisitUv  []*VisitKV `json:"visit_uv"`
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
	visit_uv = result.VisitUv
	return
}
