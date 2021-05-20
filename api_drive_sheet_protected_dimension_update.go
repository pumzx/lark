// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateSheetProtectedDimension 该接口用于根据保护范围ID修改保护范围，单次最多支持同时修改10个ID。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUTM5YjL1ETO24SNxkjN
func (r *DriveService) UpdateSheetProtectedDimension(ctx context.Context, request *UpdateSheetProtectedDimensionReq, options ...MethodOptionFunc) (*UpdateSheetProtectedDimensionResp, *Response, error) {
	if r.cli.mock.mockDriveUpdateSheetProtectedDimension != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#UpdateSheetProtectedDimension mock enable")
		return r.cli.mock.mockDriveUpdateSheetProtectedDimension(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "UpdateSheetProtectedDimension",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/protected_range_batch_update",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateSheetProtectedDimensionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveUpdateSheetProtectedDimension(f func(ctx context.Context, request *UpdateSheetProtectedDimensionReq, options ...MethodOptionFunc) (*UpdateSheetProtectedDimensionResp, *Response, error)) {
	r.mockDriveUpdateSheetProtectedDimension = f
}

func (r *Mock) UnMockDriveUpdateSheetProtectedDimension() {
	r.mockDriveUpdateSheetProtectedDimension = nil
}

type UpdateSheetProtectedDimensionReq struct {
	SpreadSheetToken string                                    `path:"spreadsheetToken" json:"-"` // sheet 的 token，获取方式见 [对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN)的第 4 项
	Requests         *UpdateSheetProtectedDimensionReqRequests `json:"requests,omitempty"`        // 请求
}

type UpdateSheetProtectedDimensionReqRequests struct {
	ProtectID string                                             `json:"protectId,omitempty"` // 保护范围ID，可以通过[获取表格元数据](/ssl:ttdoc/ukTMukTMukTM/uETMzUjLxEzM14SMxMTN) 接口获取
	Dimension *UpdateSheetProtectedDimensionReqRequestsDimension `json:"dimension,omitempty"` // 行列保护信息
	Editors   *UpdateSheetProtectedDimensionReqRequestsEditors   `json:"editors,omitempty"`   // 可编辑保护范围的用户
	LockInfo  *string                                            `json:"lockInfo,omitempty"`  // 保护说明
}

type UpdateSheetProtectedDimensionReqRequestsDimension struct {
	SheetID        string `json:"sheetId,omitempty"`        // sheetId
	StartIndex     int64  `json:"startIndex,omitempty"`     // 保护行列起始下标，下标从1开始
	EndIndex       int64  `json:"endIndex,omitempty"`       // 保护行列终止下标，下标从1开始
	MajorDimension string `json:"majorDimension,omitempty"` // 保护范围ID对应的保护范围的维度，COLUMNS为保护列，ROWS为保护行
}

type UpdateSheetProtectedDimensionReqRequestsEditors struct {
	AddEditors *UpdateSheetProtectedDimensionReqRequestsEditorsAddEditors `json:"addEditors,omitempty"` // 需要增加的用户的列表，用户需要有文档的编辑权限
	DelEditors *UpdateSheetProtectedDimensionReqRequestsEditorsDelEditors `json:"delEditors,omitempty"` // 需要删除的用户的列表
}

type UpdateSheetProtectedDimensionReqRequestsEditorsAddEditors struct {
	MemberType string `json:"memberType,omitempty"` // 用户类型，支持userId,openId,unionId
	MemberID   string `json:"memberId,omitempty"`   // 用户类型对应的用户ID
}

type UpdateSheetProtectedDimensionReqRequestsEditorsDelEditors struct {
	MemberType string `json:"memberType,omitempty"` // 用户类型，支持userId,openId,unionId
	MemberID   string `json:"memberId,omitempty"`   // 用户类型对应的用户ID
}

type updateSheetProtectedDimensionResp struct {
	Code int64                              `json:"code,omitempty"`
	Msg  string                             `json:"msg,omitempty"`
	Data *UpdateSheetProtectedDimensionResp `json:"data,omitempty"`
}

type UpdateSheetProtectedDimensionResp struct {
	Replies []*UpdateSheetProtectedDimensionRespReplie `json:"replies,omitempty"` // 响应
}

type UpdateSheetProtectedDimensionRespReplie struct {
	SheetID   string                                            `json:"sheetId,omitempty"`   // sheet的id
	Dimension *UpdateSheetProtectedDimensionRespReplieDimension `json:"dimension,omitempty"` // 成功修改的保护行列信息
	Editors   *UpdateSheetProtectedDimensionRespReplieEditors   `json:"editors,omitempty"`   // 可编辑保护范围的用户
	LockInfo  string                                            `json:"lockInfo,omitempty"`  // 成功修改的保护说明
}

type UpdateSheetProtectedDimensionRespReplieDimension struct {
	SheetID        string `json:"sheetId,omitempty"`        // sheetId
	StartIndex     int64  `json:"startIndex,omitempty"`     // 保护行列起始下标，下标从1开始
	EndIndex       int64  `json:"endIndex,omitempty"`       // 保护行列终止下标，下标从1开始
	MajorDimension string `json:"majorDimension,omitempty"` // 保护范围的维度
}

type UpdateSheetProtectedDimensionRespReplieEditors struct {
	AddEditors []*UpdateSheetProtectedDimensionRespReplieEditorsAddEditor `json:"addEditors,omitempty"` // 成功增加的用户的列表
	DelEditors []*UpdateSheetProtectedDimensionRespReplieEditorsDelEditor `json:"delEditors,omitempty"` // 成功删除的用户的列表
}

type UpdateSheetProtectedDimensionRespReplieEditorsAddEditor struct {
	MemberType string `json:"memberType,omitempty"` // 用户类型
	MemberID   string `json:"memberId,omitempty"`   // 用户类型对应的用户ID
}

type UpdateSheetProtectedDimensionRespReplieEditorsDelEditor struct {
	MemberType string `json:"memberType,omitempty"` // 用户类型
	MemberID   string `json:"memberId,omitempty"`   // 用户类型对应的用户ID
}
