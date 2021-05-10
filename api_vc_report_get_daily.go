package lark

import (
	"context"
)

// GetDailyReport
//
// > 获取一段时间内每日会议的使用报告
// 支持查询最近90天内的数据，当日数据在第二日上午9点计入
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/vc-v1/report/get_daily
func (r *VCAPI) GetDailyReport(ctx context.Context, request *GetDailyReportReq) (*GetDailyReportResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/vc/v1/reports/get_daily",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getDailyReportResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("VC", "GetDailyReport", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetDailyReportReq struct {
	StartTime string `query:"start_time" json:"-"` // 开始时间（unix时间，单位sec），必选字段，示例值："1608888867"
	EndTime   string `query:"end_time" json:"-"`   // 结束时间（unix时间，单位sec），必选字段，示例值："1608888966"
}

type getDailyReportResp struct {
	Code int                 `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *GetDailyReportResp `json:"data,omitempty"` // -
}

type GetDailyReportResp struct {
	MeetingReport *GetDailyReportRespMeetingReport `json:"meeting_report,omitempty"` // 会议报告数据
}

type GetDailyReportRespMeetingReport struct {
	TotalMeetingCount     string                                      `json:"total_meeting_count,omitempty"`     // 总会议数量
	TotalMeetingDuration  string                                      `json:"total_meeting_duration,omitempty"`  // 总会议时长（单位sec）
	TotalParticipantCount string                                      `json:"total_participant_count,omitempty"` // 总参会人数
	DailyReport           *GetDailyReportRespMeetingReportDailyReport `json:"daily_report,omitempty"`            // 每日会议报告列表
}

type GetDailyReportRespMeetingReportDailyReport struct {
	Date             string `json:"date,omitempty"`              // 日期（unix时间，单位sec）
	MeetingCount     string `json:"meeting_count,omitempty"`     // 会议数量
	MeetingDuration  string `json:"meeting_duration,omitempty"`  // 会议时长（单位sec）
	ParticipantCount string `json:"participant_count,omitempty"` // 参会人数
}
