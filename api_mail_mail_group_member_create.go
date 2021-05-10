package lark

import (
	"context"
)

// CreateMailGroupMember 向邮件组添加单个成员
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-member/create
func (r *MailAPI) CreateMailGroupMember(ctx context.Context, request *CreateMailGroupMemberReq) (*CreateMailGroupMemberResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id/members",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(createMailGroupMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Mail", "CreateMailGroupMember", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type CreateMailGroupMemberReq struct {
	UserIDType       *IDType           `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门ID的类型, 示例值："open_department_id", 可选值有: `department_id`：以自定义department_id来标识部门, `open_department_id`：以open_department_id来标识部门
	MailGroupID      string            `path:"mailgroup_id" json:"-"`        // 邮件组ID或者邮件组地址, 示例值："xxxxxxxxxxxxxxx 或 test_mail_group@xxx.xx"
	Email            *string           `json:"email,omitempty"`              // 成员邮箱地址（当成员类型是EXTERNAL_USER/MAIL_GROUP/OTHER_MEMBER时有值）, 示例值："test_memeber@xxx.xx"
	UserID           *string           `json:"user_id,omitempty"`            // 租户内用户的唯一标识（当成员类型是USER时有值）, 示例值："xxxxxxxxxx"
	DepartmentID     *string           `json:"department_id,omitempty"`      // 租户内部门的唯一标识（当成员类型是DEPARTMENT时有值）, 示例值："xxxxxxxxxx"
	Type             *MailUserType     `json:"type,omitempty"`               // 成员类型, 示例值："USER", 可选值有: `USER`：内部用户, `DEPARTMENT`：部门, `COMPANY`：全员, `EXTERNAL_USER`：外部用户, `MAIL_GROUP`：邮件组, `OTHER_MEMBER`：内部成员
}

type createMailGroupMemberResp struct {
	Code int                        `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *CreateMailGroupMemberResp `json:"data,omitempty"` //
}

type CreateMailGroupMemberResp struct {
	MemberID     string       `json:"member_id,omitempty"`     // 邮件组内成员唯一标识
	Email        string       `json:"email,omitempty"`         // 成员邮箱地址（当成员类型是EXTERNAL_USER/MAIL_GROUP/OTHER_MEMBER时有值）
	UserID       string       `json:"user_id,omitempty"`       // 租户内用户的唯一标识（当成员类型是USER时有值）
	DepartmentID string       `json:"department_id,omitempty"` // 租户内部门的唯一标识（当成员类型是DEPARTMENT时有值）
	Type         MailUserType `json:"type,omitempty"`          // 成员类型, 可选值有: `USER`：内部用户, `DEPARTMENT`：部门, `COMPANY`：全员, `EXTERNAL_USER`：外部用户, `MAIL_GROUP`：邮件组, `OTHER_MEMBER`：内部成员
}
