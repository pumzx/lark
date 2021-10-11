// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetChatAnnouncement 获取会话中的群公告信息，公告信息格式与[云文档](https://open.feishu.cn/document/ukTMukTMukTM/uAzM5YjLwMTO24CMzkjN)格式相同。
//
// 注意事项：
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/home/develop-a-bot-in-5-minutes/create-an-app)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-announcement/get
func (r *ChatService) GetChatAnnouncement(ctx context.Context, request *GetChatAnnouncementReq, options ...MethodOptionFunc) (*GetChatAnnouncementResp, *Response, error) {
	if r.cli.mock.mockChatGetChatAnnouncement != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#GetChatAnnouncement mock enable")
		return r.cli.mock.mockChatGetChatAnnouncement(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "GetChatAnnouncement",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats/:chat_id/announcement",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getChatAnnouncementResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockChatGetChatAnnouncement(f func(ctx context.Context, request *GetChatAnnouncementReq, options ...MethodOptionFunc) (*GetChatAnnouncementResp, *Response, error)) {
	r.mockChatGetChatAnnouncement = f
}

func (r *Mock) UnMockChatGetChatAnnouncement() {
	r.mockChatGetChatAnnouncement = nil
}

type GetChatAnnouncementReq struct {
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求:  获取用户 user ID
	ChatID     string  `path:"chat_id" json:"-"`       // 待获取公告的群 ID，详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 示例值："oc_5ad11d72b830411d72b836c20"
}

type getChatAnnouncementResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *GetChatAnnouncementResp `json:"data,omitempty"`
}

type GetChatAnnouncementResp struct {
	Content        string `json:"content,omitempty"`          // 云文档序列化信息
	Revision       string `json:"revision,omitempty"`         // 文档当前版本号 纯数字
	CreateTime     string `json:"create_time,omitempty"`      // 文档生成的时间戳（秒）
	UpdateTime     string `json:"update_time,omitempty"`      // 文档更新的时间戳（秒）
	OwnerIDType    IDType `json:"owner_id_type,omitempty"`    // 文档所有者的 ID 类型, 如果所有者是用户，则与查询参数中的user_id_type 相同；取值为`open_id` `user_id` `union_id` 其中之一，不同 ID 的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 如果所有者是机器人，为机器人应用的 `app_id`，详情参见  [获取应用身份访问凭证](https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/g), 可选值有: `user_id`：以 user_id 来识别用户, `union_id`：以 union_id 来识别用户, `open_id`：以 open_id 来识别用户, `app_id`：以 app_id 来识别机器人应用
	OwnerID        string `json:"owner_id,omitempty"`         // 文档所有者 ID，ID 值与owner_id_type 中的ID类型对应
	ModifierIDType IDType `json:"modifier_id_type,omitempty"` // 文档最新修改者 id 类型, - 如果修改者是用户，则与查询参数中的user_id_type 相同；取值为`open_id` `user_id` `union_id` 其中之一，不同 ID 的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 如果修改者是机器人，为机器人应用的 `app_id`，详情参见  [获取应用身份访问凭证](https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/g), 可选值有: `user_id`：以 user_id 来识别用户, `union_id`：以 union_id 来识别用户, `open_id`：以 open_id 来识别用户, `app_id`：以 app_id 来识别应用
	ModifierID     string `json:"modifier_id,omitempty"`      // 文档最新修改者 ID，ID 值与modifier_id_type 中的ID类型对应
}
