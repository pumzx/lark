package lark

import (
	"context"
)

// GetTopUserReport
//
// > 获取一段时间内使用会议的top用户列表
// 支持最近查询最近90天内的数据，当日数据在第二日上午9点计入；默认返回前10位，最多可查询前100位
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/vc-v1/report/get_top_user
func (r *VCAPI) GetTopUserReport(ctx context.Context, request *GetTopUserReportReq) (*GetTopUserReportResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/vc/v1/reports/get_top_user",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getTopUserReportResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("VC", "GetTopUserReport", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetTopUserReportReq struct {
	StartTime string `query:"start_time" json:"-"` // 开始时间（unix时间，单位sec），必选字段，示例值："1608888867"
	EndTime   string `query:"end_time" json:"-"`   // 结束时间（unix时间，单位sec），必选字段，示例值："1608889966"
	Limit     int    `query:"limit" json:"-"`      // 取前多少位，必选字段，示例值：10
	OrderBy   int    `query:"order_by" json:"-"`   // 排序依据（降序），必选字段，可用值：【1（会议数量），2（会议时长）】，示例值：1
}

type getTopUserReportResp struct {
	Code int                   `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *GetTopUserReportResp `json:"data,omitempty"` // -
}

type GetTopUserReportResp struct {
	TopUserReport *GetTopUserReportRespTopUserReport `json:"top_user_report,omitempty"` // top用户列表
}

type GetTopUserReportRespTopUserReport struct {
	ID              string `json:"id,omitempty"`               // 用户id
	Name            string `json:"name,omitempty"`             // 用户名
	UserType        int    `json:"user_type,omitempty"`        // 用户类型，可用值：【1（飞书用户），2（飞书会议室用户），3（文档用户），4（neo单品用户），5（neo单品游客用户），6（pstn用户），7（sip用户）】
	MeetingCount    string `json:"meeting_count,omitempty"`    // 会议数量
	MeetingDuration string `json:"meeting_duration,omitempty"` // 会议时长（单位sec）
}
