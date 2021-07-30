// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetAttendanceUserStatisticsSettings
//
// 查询日度统计或月度统计的统计设置信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/query-user-statistics-settings
func (r *AttendanceService) GetAttendanceUserStatisticsSettings(ctx context.Context, request *GetAttendanceUserStatisticsSettingsReq, options ...MethodOptionFunc) (*GetAttendanceUserStatisticsSettingsResp, *Response, error) {
	if r.cli.mock.mockAttendanceGetAttendanceUserStatisticsSettings != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#GetAttendanceUserStatisticsSettings mock enable")
		return r.cli.mock.mockAttendanceGetAttendanceUserStatisticsSettings(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "GetAttendanceUserStatisticsSettings",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/user_stats_views/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getAttendanceUserStatisticsSettingsResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockAttendanceGetAttendanceUserStatisticsSettings(f func(ctx context.Context, request *GetAttendanceUserStatisticsSettingsReq, options ...MethodOptionFunc) (*GetAttendanceUserStatisticsSettingsResp, *Response, error)) {
	r.mockAttendanceGetAttendanceUserStatisticsSettings = f
}

func (r *Mock) UnMockAttendanceGetAttendanceUserStatisticsSettings() {
	r.mockAttendanceGetAttendanceUserStatisticsSettings = nil
}

type GetAttendanceUserStatisticsSettingsReq struct {
	EmployeeType EmployeeType `query:"employee_type" json:"-"` // 用户 ID 类型, 可选值有: `employee_id`, `employee_no`
	Locale       string       `json:"locale,omitempty"`        // 语言类型, 可选值有: `en`：英文, `ja`：日文, `zh`：中文
	StatsType    string       `json:"stats_type,omitempty"`    // 统计类型,      , 可选值有: `daily`：日度统计, `month`：月度统计
}

type getAttendanceUserStatisticsSettingsResp struct {
	Code int64                                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                                   `json:"msg,omitempty"`  // 错误描述
	Data *GetAttendanceUserStatisticsSettingsResp `json:"data,omitempty"`
}

type GetAttendanceUserStatisticsSettingsResp struct {
	View *GetAttendanceUserStatisticsSettingsRespView `json:"view,omitempty"` // 统计视图
}

type GetAttendanceUserStatisticsSettingsRespView struct {
	ViewID    string                                             `json:"view_id,omitempty"`    // 视图 ID
	StatsType string                                             `json:"stats_type,omitempty"` // 统计类型, 可选值有: `daily`：日度统计, `month`：月度统计
	UserID    string                                             `json:"user_id,omitempty"`    // 用户 ID
	Items     []*GetAttendanceUserStatisticsSettingsRespViewItem `json:"items,omitempty"`      // 一级标题
}

type GetAttendanceUserStatisticsSettingsRespViewItem struct {
	Code       string                                                      `json:"code,omitempty"`        // 标题编号
	Title      string                                                      `json:"title,omitempty"`       // 标题名称
	ChildItems []*GetAttendanceUserStatisticsSettingsRespViewItemChildItem `json:"child_items,omitempty"` // 子标题
}

type GetAttendanceUserStatisticsSettingsRespViewItemChildItem struct {
	Code       string `json:"code,omitempty"`        // 标题编号
	Value      string `json:"value,omitempty"`       // 是否开启,      , 可选值有: `0`：关闭, `1`：开启
	Title      string `json:"title,omitempty"`       // 标题名称
	ColumnType int64  `json:"column_type,omitempty"` // 标题类型
	ReadOnly   bool   `json:"read_only,omitempty"`   // 是否只读
	MinValue   string `json:"min_value,omitempty"`   // 最小值
	MaxValue   string `json:"max_value,omitempty"`   // 最大值
}
