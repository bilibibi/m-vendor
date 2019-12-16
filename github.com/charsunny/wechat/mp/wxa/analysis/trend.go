package analysis

import (
	"github.com/charsunny/wechat/mp/core"
)

type TrendKV struct {
	RefDate string `json:"ref_date"`	// 日期，格式为 yyyymmdd
	SessionCnt int `json:"session_cnt"`	// 打开次数
	VisitPv int `json:"visit_pv"` // 访问次数
	VisitUv int `json:"visit_uv"`	// 访问人数
	VisitUvNew int `json:"visit_uv_new"`	// 新用户数
	StayTimeUv float64 `json:"stay_time_uv"`	// 人均停留时长 (浮点型，单位：秒)
	StayTimeSession float64 `json:"stay_time_session"`	// 次均停留时长 (浮点型，单位：秒)
	VisitDepth float64 `json:"visit_depth"`	// 平均访问深度 (浮点型)
}

// 获取用户访问小程序日留存
// @param begin_date	string		是	开始日期。格式为 yyyymmdd
// @param end_date	string		是	开始日期。格式为 结束日期，限定查询1天数据
// @return visitUvNew	Object	新增用户留存
// @return visitUv	Object	新增用户留存
func GetDailyTrend(clt *core.Client, begin_date string, end_date string) (list []*TrendKV, err error) {
	const incompleteURL = "https://api.weixin.qq.com/datacube/getweanalysisappiddailyvisittrend?access_token="

	var req = struct {
		BeginDate string `json:"begin_date"`
		EndDate string `json:"end_date"`
	} {
		BeginDate:begin_date,
		EndDate:end_date,
	}
	var result struct {
		core.Error
		List []*TrendKV `json:"list"`
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

// 获取用户访问小程序周留存
// @param begin_date	string		是	开始日期，为周一日期。格式为 yyyymmdd
// @param end_date	string		是	开始日期。结束日期，为周日日期，限定查询一周数据。格式为 yyyymmdd
// @return visitUvNew	Object	新增用户留存
// @return visitUv	Object	新增用户留存
func GetWeeklyTrend(clt *core.Client, begin_date string, end_date string) (list []*TrendKV, err error) {
	const incompleteURL = "https://api.weixin.qq.com/datacube/getweanalysisappidweeklyvisittrend?access_token="

	var req = struct {
		BeginDate string `json:"begin_date"`
		EndDate string `json:"end_date"`
	} {
		BeginDate:begin_date,
		EndDate:end_date,
	}
	var result struct {
		core.Error
		List []*TrendKV `json:"list"`
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

// 获取用户访问小程序月留存
// @param begin_date	string		是	开始日期，为自然月第一天。格式为 yyyymmdd
// @param end_date	string		是	开始日期。结束日期，为自然月最后一天，限定查询一个月数据。格式为 yyyymmdd
// @return visitUvNew	Object	新增用户留存
// @return visitUv	Object	新增用户留存
func GetMonthlyTrend(clt *core.Client, begin_date string, end_date string) (list []*TrendKV, err error) {
	const incompleteURL = "https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyvisittrend?access_token="

	var req = struct {
		BeginDate string `json:"begin_date"`
		EndDate string `json:"end_date"`
	} {
		BeginDate:begin_date,
		EndDate:end_date,
	}
	var result struct {
		core.Error
		List []*TrendKV `json:"list"`
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
