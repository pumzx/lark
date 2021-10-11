// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateChatAnnouncement 更新会话中的群公告信息，更新公告信息的格式和更新[云文档](https://open.feishu.cn/document/ukTMukTMukTM/uAzM5YjLwMTO24CMzkjN)格式相同。
//
// 注意事项：
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/home/develop-a-bot-in-5-minutes/create-an-app)
// - 当授权用户或机器人是群主时，可更新群公告信息
// - 当授权用户或机器人非群主，且群主未设置 [仅群主可编辑群信息] 时，可更新群公告信息
// - 当授权用户或机器人非群主，且群主设置了 [仅群主可编辑群信息] 时，无法更新公告信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-announcement/patch
func (r *ChatService) UpdateChatAnnouncement(ctx context.Context, request *UpdateChatAnnouncementReq, options ...MethodOptionFunc) (*UpdateChatAnnouncementResp, *Response, error) {
	if r.cli.mock.mockChatUpdateChatAnnouncement != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#UpdateChatAnnouncement mock enable")
		return r.cli.mock.mockChatUpdateChatAnnouncement(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "UpdateChatAnnouncement",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats/:chat_id/announcement",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateChatAnnouncementResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockChatUpdateChatAnnouncement(f func(ctx context.Context, request *UpdateChatAnnouncementReq, options ...MethodOptionFunc) (*UpdateChatAnnouncementResp, *Response, error)) {
	r.mockChatUpdateChatAnnouncement = f
}

func (r *Mock) UnMockChatUpdateChatAnnouncement() {
	r.mockChatUpdateChatAnnouncement = nil
}

type UpdateChatAnnouncementReq struct {
	ChatID   string   `path:"chat_id" json:"-"`   // 待修改公告的群 ID，详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 示例值："oc_5ad11d72b830411d72b836c20"
	Revision string   `json:"revision,omitempty"` // 文档当前版本号 int64 类型，get 接口会返回, 示例值："12"
	Requests []string `json:"requests,omitempty"` // 修改文档请求的序列化字段, 示例值：xxx
}

type updateChatAnnouncementResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *UpdateChatAnnouncementResp `json:"data,omitempty"`
}

type UpdateChatAnnouncementResp struct{}
