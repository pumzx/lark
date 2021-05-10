package lark

import (
	"context"
)

// ApplyReserve
//
// > 预约一个会议
// 支持预约到期时间在未来30天内的会议，预约到期后会议号将被释放，如需继续使用可通过"更新预约"接口进行续期；如有会议权限的特殊需求，可在预约时进行配置
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/vc-v1/reserve/apply
func (r *VCAPI) ApplyReserve(ctx context.Context, request *ApplyReserveReq) (*ApplyReserveResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/vc/v1/reserves/apply",
		Body:                  request,
		NeedTenantAccessToken: false,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(applyReserveResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("VC", "ApplyReserve", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type ApplyReserveReq struct {
	UserIDType      IDType                          `query:"user_id_type" json:"-"`     // 用户ID类型，可用值：【open_id，union_id，user_id】，默认值：open_id
	EndTime         string                          `json:"end_time,omitempty"`         // 预约到期时间（unix时间，单位sec），必选字段
	MeetingSettings *ApplyReserveReqMeetingSettings `json:"meeting_settings,omitempty"` // 会议设置
}

type ApplyReserveReqMeetingSettings struct {
	Topic             string                                            `json:"topic,omitempty"`              // 会议主题
	ActionPermissions []*ApplyReserveReqMeetingSettingsActionPermission `json:"action_permissions,omitempty"` // 会议权限配置列表，示例请求体中的权限配置含义为：配置"是否能成为主持人"的权限，检查用户id，检查方式为白名单，即：在白名单中的用户才有成为主持人的权限。如没有限制权限的需求，该字段可为空；如列表中存在相同的权限项，则它们之间为"逻辑或"的关系（即 有一个为true则拥有该权限）
}

type ApplyReserveReqMeetingSettingsActionPermission struct {
	Permission         int                                                                `json:"permission,omitempty"`          // 权限项，必选字段，可用值：【1（是否能成为主持人），2（是否能邀请参会人），3（是否能分享会议）】
	PermissionCheckers []*ApplyReserveReqMeetingSettingsActionPermissionPermissionChecker `json:"permission_checkers,omitempty"` // 权限检查器列表，同一项权限可配置多个权限检查器，权限检查器之间为"逻辑或"的关系（即 有一个为true则结果为true），必选字段
}

type ApplyReserveReqMeetingSettingsActionPermissionPermissionChecker struct {
	CheckField int      `json:"check_field,omitempty"` // 检查字段类型，即check_list中的值的类型，如用户id、租户id等，必选字段，可用值：【1（用户id），2（用户类型），3（租户ID）】
	CheckMode  int      `json:"check_mode,omitempty"`  // 检查方式，白名单或者黑名单，必选字段，可用值：【1（白名单（在check_list中为有权限）），2（黑名单（不在check_list中为有权限））】
	CheckList  []string `json:"check_list,omitempty"`  // 检查字段值列表，根据check_mode不同为白名单或黑名单，必选字段
}

type applyReserveResp struct {
	Code int               `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string            `json:"msg,omitempty"`  // 错误描述
	Data *ApplyReserveResp `json:"data,omitempty"` // -
}

type ApplyReserveResp struct {
	Reserve *ApplyReserveRespReserve `json:"reserve,omitempty"` // 预约数据
}

type ApplyReserveRespReserve struct {
	ID        string `json:"id,omitempty"`         // 预约id
	MeetingNo string `json:"meeting_no,omitempty"` // 9位会议号
	URL       string `json:"url,omitempty"`        // 会议链接
	AppLink   string `json:"app_link,omitempty"`   // APPLink用于唤起飞书APP入会(目前仅支持飞书移动端)，模板见下面的示例。其中"{?}"为占位符，用于配置入会参数，使用需替换为具体配置值：preview为入会前的设置页，mic为麦克风，speaker为扬声器，camera为摄像头，配置0表示关闭，配置1表示打开。例如："https://applink.feishu.cn/client/videochat/open?source=openplatform&action=join&idtype=reservationid&id=6911654912696467457&preview=0&mic=1&speaker=1&camera=0"
	EndTime   string `json:"end_time,omitempty"`   // 预约到期时间（unix时间，单位sec）
}
