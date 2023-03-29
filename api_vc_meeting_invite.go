// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// InviteVCMeeting 邀请参会人进入会议。
//
// 发起邀请的操作者必须具有相应的权限（如果操作者为用户, 则必须在会中）, 如果会议被锁定、或参会人数如果达到上限, 则会邀请失败
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/invite
func (r *VCService) InviteVCMeeting(ctx context.Context, request *InviteVCMeetingReq, options ...MethodOptionFunc) (*InviteVCMeetingResp, *Response, error) {
	if r.cli.mock.mockVCInviteVCMeeting != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#InviteVCMeeting mock enable")
		return r.cli.mock.mockVCInviteVCMeeting(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "VC",
		API:                 "InviteVCMeeting",
		Method:              "PATCH",
		URL:                 r.cli.openBaseURL + "/open-apis/vc/v1/meetings/:meeting_id/invite",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(inviteVCMeetingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCInviteVCMeeting mock VCInviteVCMeeting method
func (r *Mock) MockVCInviteVCMeeting(f func(ctx context.Context, request *InviteVCMeetingReq, options ...MethodOptionFunc) (*InviteVCMeetingResp, *Response, error)) {
	r.mockVCInviteVCMeeting = f
}

// UnMockVCInviteVCMeeting un-mock VCInviteVCMeeting method
func (r *Mock) UnMockVCInviteVCMeeting() {
	r.mockVCInviteVCMeeting = nil
}

// InviteVCMeetingReq ...
type InviteVCMeetingReq struct {
	MeetingID  string                       `path:"meeting_id" json:"-"`    // 会议ID（视频会议的唯一标识, 视频会议开始后才会产生）, 示例值: "6911188411932033028"
	UserIDType *IDType                      `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	Invitees   []*InviteVCMeetingReqInvitee `json:"invitees,omitempty"`     // 被邀请的用户列表【一次性最多支持邀请10人】
}

// InviteVCMeetingReqInvitee ...
type InviteVCMeetingReqInvitee struct {
	ID       *string `json:"id,omitempty"`        // 用户ID, 示例值: "ou_3ec3f6a28a0d08c45d895276e8e5e19b"
	UserType *int64  `json:"user_type,omitempty"` // 用户类型, 示例值: 1, 可选值有: 1: 飞书用户, 2: rooms用户, 3: 文档用户, 4: neo单品用户, 5: neo单品游客用户, 6: pstn用户, 7: sip用户
}

// InviteVCMeetingResp ...
type InviteVCMeetingResp struct {
	InviteResults []*InviteVCMeetingRespInviteResult `json:"invite_results,omitempty"` // 邀请结果
}

// InviteVCMeetingRespInviteResult ...
type InviteVCMeetingRespInviteResult struct {
	ID       string `json:"id,omitempty"`        // 用户ID
	UserType int64  `json:"user_type,omitempty"` // 用户类型, 可选值有: 1: 飞书用户, 2: rooms用户, 3: 文档用户, 4: neo单品用户, 5: neo单品游客用户, 6: pstn用户, 7: sip用户
	Status   int64  `json:"status,omitempty"`    // 邀请结果, 可选值有: 1: 邀请成功, 2: 邀请失败
}

// inviteVCMeetingResp ...
type inviteVCMeetingResp struct {
	Code int64                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *InviteVCMeetingResp `json:"data,omitempty"`
}
