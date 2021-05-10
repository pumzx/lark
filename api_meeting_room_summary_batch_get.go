package lark

import (
	"context"
)

// BatchGetSummary 通过日程的Uid和Original time，查询会议室日程主题。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uIjM5UjLyITO14iMykTN/
func (r *MeetingRoomAPI) BatchGetSummary(ctx context.Context, request *BatchGetSummaryReq) (*BatchGetSummaryResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/summary/batch_get",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(batchGetSummaryResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("MeetingRoom", "BatchGetSummary", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type BatchGetSummaryReq struct {
	EventUids *BatchGetSummaryReqEventUids `json:"EventUids,omitempty"` // 需要查询的日程Uid和Original time
}

type BatchGetSummaryReqEventUids struct {
	Uid          string `json:"uid,omitempty"`           // 日程的唯一id
	OriginalTime int    `json:"original_time,omitempty"` // 日程实例原始时间，非重复日程必为0。重复日程若为0则表示回复其所有实例，否则表示回复单个实例。
}

type batchGetSummaryResp struct {
	Code int                  `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *BatchGetSummaryResp `json:"data,omitempty"` // 返回业务信息
}

type BatchGetSummaryResp struct {
	EventInfos     *BatchGetSummaryRespEventInfos     `json:"EventInfos,omitempty"`     // 成功查询到的日程信息
	ErrorEventUids *BatchGetSummaryRespErrorEventUids `json:"ErrorEventUids,omitempty"` // 没有查询到的日程
}

type BatchGetSummaryRespEventInfos struct {
	Uid          string `json:"uid,omitempty"`           // 日程的唯一id
	OriginalTime int    `json:"original_time,omitempty"` // 日程实例原始时间，非重复日程必为0。重复日程若为0则表示回复其所有实例，否则表示回复单个实例。
	Summary      string `json:"summary,omitempty"`       // 日程主题
}

type BatchGetSummaryRespErrorEventUids struct {
	Uid          string `json:"uid,omitempty"`           // 日程的唯一id
	OriginalTime int    `json:"original_time,omitempty"` // 日程实例原始时间，非重复日程必为0。重复日程若为0则表示回复其所有实例，否则表示回复单个实例。
	ErrorMsg     string `json:"error_msg,omitempty"`     // 错误信息
}
