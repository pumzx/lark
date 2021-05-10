package lark

import (
	"context"
)

// GetDistrictList 新建建筑时需要选择所处国家/地区，该接口用于获得系统预先提供的可供选择的城市列表。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUTNwYjL1UDM24SN1AjN
func (r *MeetingRoomAPI) GetDistrictList(ctx context.Context, request *GetDistrictListReq) (*GetDistrictListResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/district/list?country_id=1814991",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getDistrictListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("MeetingRoom", "GetDistrictList", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetDistrictListReq struct {
	CountryID string `query:"country_id" json:"-"` // 国家地区ID
}

type getDistrictListResp struct {
	Code int                  `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *GetDistrictListResp `json:"data,omitempty"` // 返回业务信息
}

type GetDistrictListResp struct {
	Districts *GetDistrictListRespDistricts `json:"districts,omitempty"` // 城市列表
}

type GetDistrictListRespDistricts struct {
	DistrictID string `json:"district_id,omitempty"` // 城市ID
	Name       string `json:"name,omitempty"`        // 城市名称
}
