package lark

import (
	"context"
)

// BatchGetBuilding 该接口用于获取指定建筑物的详细信息。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ukzNyUjL5cjM14SO3ITN
func (r *MeetingRoomAPI) BatchGetBuilding(ctx context.Context, request *BatchGetBuildingReq) (*BatchGetBuildingResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/building/batch_get?building_ids=omb_8ec170b937536a5d87c23b418b83f9bb&building_ids=omb_38570e4f0fd9ecf15030d3cc8b388f3a&fields=*",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(batchGetBuildingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("MeetingRoom", "BatchGetBuilding", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type BatchGetBuildingReq struct {
	BuildingIDs string  `query:"building_ids" json:"-"` // 用于查询指定建筑物的 ID
	Fields      *string `query:"fields" json:"-"`       // 用于指定返回的字段名，每个字段名之间用逗号 "," 分隔，如：“id,name”，"*" 表示返回全部字段，可选字段有："id,name,description,floors"，默认返回所有字段
}

type batchGetBuildingResp struct {
	Code int                   `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *BatchGetBuildingResp `json:"data,omitempty"` // 返回业务信息
}

type BatchGetBuildingResp struct {
	Buildings *BatchGetBuildingRespBuildings `json:"buildings,omitempty"` // 建筑列表
}

type BatchGetBuildingRespBuildings struct {
	BuildingID  string   `json:"building_id,omitempty"` // 建筑物 ID
	Description string   `json:"description,omitempty"` // 建筑物的相关描述
	Floors      []string `json:"floors,omitempty"`      // 属于当前建筑物的所有楼层列表
	Name        string   `json:"name,omitempty"`        // 建筑物名称
}
