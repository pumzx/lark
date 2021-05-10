package lark

import (
	"context"
)

// CreateBuilding 该接口对应管理后台的添加建筑，添加楼层的功能，可用于创建建筑物和建筑物的楼层信息。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uATNwYjLwUDM24CM1AjN
func (r *MeetingRoomAPI) CreateBuilding(ctx context.Context, request *CreateBuildingReq) (*CreateBuildingResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/building/create",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(createBuildingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("MeetingRoom", "CreateBuilding", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type CreateBuildingReq struct {
	Name             string   `json:"name,omitempty"`               // 建筑名称
	Floors           []string `json:"floors,omitempty"`             // 楼层列表
	CountryID        string   `json:"country_id,omitempty"`         // 国家/地区ID
	DistrictID       string   `json:"district_id,omitempty"`        // 城市ID
	CustomBuildingID *string  `json:"custom_building_id,omitempty"` // 租户自定义建筑ID
}

type createBuildingResp struct {
	Code int                 `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *CreateBuildingResp `json:"data,omitempty"` // 返回业务信息
}

type CreateBuildingResp struct {
	BuildingID string `json:"building_id,omitempty"` // 成功创建的建筑ID
}
