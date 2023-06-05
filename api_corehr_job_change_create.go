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

// CreateCoreHrJobChange 创建员工异动信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/job_change/create
func (r *CoreHrService) CreateCoreHrJobChange(ctx context.Context, request *CreateCoreHrJobChangeReq, options ...MethodOptionFunc) (*CreateCoreHrJobChangeResp, *Response, error) {
	if r.cli.mock.mockCoreHrCreateCoreHrJobChange != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHr#CreateCoreHrJobChange mock enable")
		return r.cli.mock.mockCoreHrCreateCoreHrJobChange(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHr",
		API:                   "CreateCoreHrJobChange",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/job_changes",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createCoreHrJobChangeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHrCreateCoreHrJobChange mock CoreHrCreateCoreHrJobChange method
func (r *Mock) MockCoreHrCreateCoreHrJobChange(f func(ctx context.Context, request *CreateCoreHrJobChangeReq, options ...MethodOptionFunc) (*CreateCoreHrJobChangeResp, *Response, error)) {
	r.mockCoreHrCreateCoreHrJobChange = f
}

// UnMockCoreHrCreateCoreHrJobChange un-mock CoreHrCreateCoreHrJobChange method
func (r *Mock) UnMockCoreHrCreateCoreHrJobChange() {
	r.mockCoreHrCreateCoreHrJobChange = nil
}

// CreateCoreHrJobChangeReq ...
type CreateCoreHrJobChangeReq struct {
	UserIDType                   *IDType                               `query:"user_id_type" json:"-"`                    // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_admin_id: 以people_admin_id来识别用户, people_corehr_id: 以飞书人事的ID来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType             *DepartmentIDType                     `query:"department_id_type" json:"-"`              // 此次调用中使用的部门 ID 类型, 示例值: people_corehr_department_id, 可选值有: open_department_id: 以 open_department_id 来标识部门, department_id: 以 department_id 来标识部门, people_corehr_department_id: 以 people_corehr_department_id 来标识部门, 默认值: `people_corehr_department_id`
	TransferMode                 int64                                 `json:"transfer_mode,omitempty"`                   // 异动方式, 示例值: 2, 可选值有: 1: 直接异动, 2: 发起异动
	EmploymentID                 string                                `json:"employment_id,omitempty"`                   // 雇员id, 示例值: "ou_a294793e8fa21529f2a60e3e9de45520"
	TransferTypeUniqueIdentifier string                                `json:"transfer_type_unique_identifier,omitempty"` // 异动类型唯一标识, 示例值: "internal_transfer"
	FlowID                       *string                               `json:"flow_id,omitempty"`                         // 关联流程唯一标识符, 可通过「获取异动类型列表」接口获得, 示例值: "people_6963913041981490725_6983885526583627531"
	EffectiveDate                string                                `json:"effective_date,omitempty"`                  // 生效日期, 示例值: "2022-03-01"
	TransferInfo                 *CreateCoreHrJobChangeReqTransferInfo `json:"transfer_info,omitempty"`                   // 异动详细信息
	TransferKey                  *string                               `json:"transfer_key,omitempty"`                    // 异动记录标识符, 示例值: "transfer_3627531"
	InitiatorID                  *string                               `json:"initiator_id,omitempty"`                    // 异动发起人 ID, 示例值: "ou_a294793e8fa21529f2a60e3e9de45520"
}

// CreateCoreHrJobChangeReqTransferInfo ...
type CreateCoreHrJobChangeReqTransferInfo struct {
	Remark                     *string                                                       `json:"remark,omitempty"`                        // 备注, 示例值: "异动详情"
	OfferInfo                  *string                                                       `json:"offer_info,omitempty"`                    // offer信息, 示例值: "优质人才, 加急处理"
	TargetDottedManagerClean   *bool                                                         `json:"target_dotted_manager_clean,omitempty"`   // 是否撤销虚线上级, 示例值: true
	ProbationExist             *bool                                                         `json:"probation_exist,omitempty"`               // 是否有试用期, 示例值: false
	OriginalDepartment         *string                                                       `json:"original_department,omitempty"`           // 原部门, 示例值: "6966236933198579208"
	TargetDepartment           *string                                                       `json:"target_department,omitempty"`             // 新部门, 示例值: "6966236933198579208"
	OriginalWorkLocation       *string                                                       `json:"original_work_location,omitempty"`        // 原工作地点, 示例值: "6967271100992587295"
	TargetWorkLocation         *string                                                       `json:"target_work_location,omitempty"`          // 新工作地点, 示例值: "6967271100992587295"
	OriginalDirectManager      *string                                                       `json:"original_direct_manager,omitempty"`       // 原直属上级, 示例值: "6974641477444060708"
	TargetDirectManager        *string                                                       `json:"target_direct_manager,omitempty"`         // 新直属上级, 示例值: "7013619729281713671"
	OriginalDottedManager      *string                                                       `json:"original_dotted_manager,omitempty"`       // 原虚线上级, 示例值: "6974648866876573198"
	TargetDottedManager        *string                                                       `json:"target_dotted_manager,omitempty"`         // 新虚线上级, 示例值: "7013328578351842852"
	OriginalJob                *string                                                       `json:"original_job,omitempty"`                  // 原职务, 示例值: "6969469398088287751"
	TargetJob                  *string                                                       `json:"target_job,omitempty"`                    // 新职务, 示例值: "6969469557836760606"
	OriginalJobFamily          *string                                                       `json:"original_job_family,omitempty"`           // 原序列, 示例值: "6967287547462419975"
	TargetJobFamily            *string                                                       `json:"target_job_family,omitempty"`             // 新序列, 示例值: "6967287547462419975"
	OriginalJobLevel           *string                                                       `json:"original_job_level,omitempty"`            // 原级别, 示例值: "6972085707674355214"
	TargetJobLevel             *string                                                       `json:"target_job_level,omitempty"`              // 新级别, 示例值: "6972085707674355214"
	OriginalWorkforceType      *string                                                       `json:"original_workforce_type,omitempty"`       // 原人员类型, 示例值: "6968386026792289828"
	TargetWorkforceType        *string                                                       `json:"target_workforce_type,omitempty"`         // 新人员类型, 示例值: "7036268995372303885"
	OriginalCompany            *string                                                       `json:"original_company,omitempty"`              // 原公司, 示例值: "6974659700705068581"
	TargetCompany              *string                                                       `json:"target_company,omitempty"`                // 新公司, 示例值: "6974659700705068581"
	OriginalContractNumber     *string                                                       `json:"original_contract_number,omitempty"`      // 原合同编号, 示例值: "55332"
	TargetContractNumber       *string                                                       `json:"target_contract_number,omitempty"`        // 新合同编号, 示例值: "55333"
	OriginalContractType       *string                                                       `json:"original_contract_type,omitempty"`        // 原合同类型, 示例值: "labor_contract"
	TargetContractType         *string                                                       `json:"target_contract_type,omitempty"`          // 新合同类型, 示例值: "labor_contract"
	OriginalDurationType       *string                                                       `json:"original_duration_type,omitempty"`        // 原期限类型, 示例值: "fixed_term"
	TargetDurationType         *string                                                       `json:"target_duration_type,omitempty"`          // 新期限类型, 示例值: "fixed_term"
	OriginalSigningType        *string                                                       `json:"original_signing_type,omitempty"`         // 原签订类型, 示例值: "new"
	TargetSigningType          *string                                                       `json:"target_signing_type,omitempty"`           // 新签订类型, 示例值: "new"
	OriginalContractStartDate  *string                                                       `json:"original_contract_start_date,omitempty"`  // 原合同开始日期, 示例值: "2021-07-01"
	TargetContractStartDate    *string                                                       `json:"target_contract_start_date,omitempty"`    // 新合同开始日期, 示例值: "2021-07-01"
	OriginalContractEndDate    *string                                                       `json:"original_contract_end_date,omitempty"`    // 原合同结束日期, 示例值: "2024-07-01"
	TargetContractEndDate      *string                                                       `json:"target_contract_end_date,omitempty"`      // 新合同结束日期, 示例值: "2024-07-01"
	OriginalWorkingHoursType   *string                                                       `json:"original_working_hours_type,omitempty"`   // 原工时制度, 示例值: "6969087376740206087"
	TargetWorkingHoursType     *string                                                       `json:"target_working_hours_type,omitempty"`     // 新工时制度, 示例值: "6969087376740206087"
	OriginalWorkingCalendar    *string                                                       `json:"original_working_calendar,omitempty"`     // 原工作日历, 示例值: "6969087376740236087"
	TargetWorkingCalendar      *string                                                       `json:"target_working_calendar,omitempty"`       // 新工作日历, 示例值: "6969087376740236087"
	OriginalProbationEndDate   *string                                                       `json:"original_probation_end_date,omitempty"`   // 原试用期预计结束日期, 示例值: "2021-11-17"
	TargetProbationEndDate     *string                                                       `json:"target_probation_end_date,omitempty"`     // 新试用期预计结束日期, 示例值: "2021-11-17"
	OriginalWeeklyWorkingHours *string                                                       `json:"original_weekly_working_hours,omitempty"` // 原周工作时长, 示例值: "162"
	TargetWeeklyWorkingHours   *string                                                       `json:"target_weekly_working_hours,omitempty"`   // 新周工作时长, 示例值: "160"
	OriginalWorkShift          *string                                                       `json:"original_work_shift,omitempty"`           // 原排班, 示例值: "work_shift"
	TargetWorkShift            *string                                                       `json:"target_work_shift,omitempty"`             // 新排班, 示例值: "non_work_shift"
	OriginalCostCenterRate     []*CreateCoreHrJobChangeReqTransferInfoOriginalCostCenterRate `json:"original_cost_center_rate,omitempty"`     // 原成本中心分摊信息
	TargetCostCenterRate       []*CreateCoreHrJobChangeReqTransferInfoTargetCostCenterRate   `json:"target_cost_center_rate,omitempty"`       // 新成本中心分摊信息
}

// CreateCoreHrJobChangeReqTransferInfoOriginalCostCenterRate ...
type CreateCoreHrJobChangeReqTransferInfoOriginalCostCenterRate struct {
	CostCenterID *string `json:"cost_center_id,omitempty"` // 支持的成本中心id, 示例值: "6950635856373745165"
	Rate         *int64  `json:"rate,omitempty"`           // 分摊比例, 示例值: 100
}

// CreateCoreHrJobChangeReqTransferInfoTargetCostCenterRate ...
type CreateCoreHrJobChangeReqTransferInfoTargetCostCenterRate struct {
	CostCenterID *string `json:"cost_center_id,omitempty"` // 支持的成本中心id, 示例值: "6950635856373745165"
	Rate         *int64  `json:"rate,omitempty"`           // 分摊比例, 示例值: 100
}

// CreateCoreHrJobChangeResp ...
type CreateCoreHrJobChangeResp struct {
	JobChangeID                    string                                 `json:"job_change_id,omitempty"`                     // 异动记录 id
	EmploymentID                   string                                 `json:"employment_id,omitempty"`                     // 雇员 id
	Status                         string                                 `json:"status,omitempty"`                            // 异动状态, 可选值有: 0: Approving  审批中, 1: Approved  审批通过, 2: Transformed  已异动, 3: Rejected  已拒绝, 4: Cancelled  已撤销
	TransferTypeUniqueIdentifier   string                                 `json:"transfer_type_unique_identifier,omitempty"`   // 异动类型唯一标识
	TransferReasonUniqueIdentifier string                                 `json:"transfer_reason_unique_identifier,omitempty"` // 异动原因唯一标识
	ProcessID                      string                                 `json:"process_id,omitempty"`                        // 异动发起后审批流程 id
	EffectiveDate                  string                                 `json:"effective_date,omitempty"`                    // 异动生效日期
	CreatedTime                    string                                 `json:"created_time,omitempty"`                      // 创建时间
	TransferInfo                   *CreateCoreHrJobChangeRespTransferInfo `json:"transfer_info,omitempty"`                     // 异动详细信息
}

// CreateCoreHrJobChangeRespTransferInfo ...
type CreateCoreHrJobChangeRespTransferInfo struct {
	Remark                     string                                                         `json:"remark,omitempty"`                        // 备注
	OfferInfo                  string                                                         `json:"offer_info,omitempty"`                    // offer信息
	TargetDottedManagerClean   bool                                                           `json:"target_dotted_manager_clean,omitempty"`   // 是否撤销虚线上级
	ProbationExist             bool                                                           `json:"probation_exist,omitempty"`               // 是否有试用期
	OriginalDepartment         string                                                         `json:"original_department,omitempty"`           // 原部门
	TargetDepartment           string                                                         `json:"target_department,omitempty"`             // 新部门
	OriginalWorkLocation       string                                                         `json:"original_work_location,omitempty"`        // 原工作地点
	TargetWorkLocation         string                                                         `json:"target_work_location,omitempty"`          // 新工作地点
	OriginalDirectManager      string                                                         `json:"original_direct_manager,omitempty"`       // 原直属上级
	TargetDirectManager        string                                                         `json:"target_direct_manager,omitempty"`         // 新直属上级
	OriginalDottedManager      string                                                         `json:"original_dotted_manager,omitempty"`       // 原虚线上级
	TargetDottedManager        string                                                         `json:"target_dotted_manager,omitempty"`         // 新虚线上级
	OriginalJob                string                                                         `json:"original_job,omitempty"`                  // 原职务
	TargetJob                  string                                                         `json:"target_job,omitempty"`                    // 新职务
	OriginalJobFamily          string                                                         `json:"original_job_family,omitempty"`           // 原序列
	TargetJobFamily            string                                                         `json:"target_job_family,omitempty"`             // 新序列
	OriginalJobLevel           string                                                         `json:"original_job_level,omitempty"`            // 原级别
	TargetJobLevel             string                                                         `json:"target_job_level,omitempty"`              // 新级别
	OriginalWorkforceType      string                                                         `json:"original_workforce_type,omitempty"`       // 原人员类型
	TargetWorkforceType        string                                                         `json:"target_workforce_type,omitempty"`         // 新人员类型
	OriginalCompany            string                                                         `json:"original_company,omitempty"`              // 原公司
	TargetCompany              string                                                         `json:"target_company,omitempty"`                // 新公司
	OriginalContractNumber     string                                                         `json:"original_contract_number,omitempty"`      // 原合同编号
	TargetContractNumber       string                                                         `json:"target_contract_number,omitempty"`        // 新合同编号
	OriginalContractType       string                                                         `json:"original_contract_type,omitempty"`        // 原合同类型
	TargetContractType         string                                                         `json:"target_contract_type,omitempty"`          // 新合同类型
	OriginalDurationType       string                                                         `json:"original_duration_type,omitempty"`        // 原期限类型
	TargetDurationType         string                                                         `json:"target_duration_type,omitempty"`          // 新期限类型
	OriginalSigningType        string                                                         `json:"original_signing_type,omitempty"`         // 原签订类型
	TargetSigningType          string                                                         `json:"target_signing_type,omitempty"`           // 新签订类型
	OriginalContractStartDate  string                                                         `json:"original_contract_start_date,omitempty"`  // 原合同开始日期
	TargetContractStartDate    string                                                         `json:"target_contract_start_date,omitempty"`    // 新合同开始日期
	OriginalContractEndDate    string                                                         `json:"original_contract_end_date,omitempty"`    // 原合同结束日期
	TargetContractEndDate      string                                                         `json:"target_contract_end_date,omitempty"`      // 新合同结束日期
	OriginalWorkingHoursType   string                                                         `json:"original_working_hours_type,omitempty"`   // 原工时制度
	TargetWorkingHoursType     string                                                         `json:"target_working_hours_type,omitempty"`     // 新工时制度
	OriginalWorkingCalendar    string                                                         `json:"original_working_calendar,omitempty"`     // 原工作日历
	TargetWorkingCalendar      string                                                         `json:"target_working_calendar,omitempty"`       // 新工作日历
	OriginalProbationEndDate   string                                                         `json:"original_probation_end_date,omitempty"`   // 原试用期预计结束日期
	TargetProbationEndDate     string                                                         `json:"target_probation_end_date,omitempty"`     // 新试用期预计结束日期
	OriginalWeeklyWorkingHours string                                                         `json:"original_weekly_working_hours,omitempty"` // 原周工作时长
	TargetWeeklyWorkingHours   string                                                         `json:"target_weekly_working_hours,omitempty"`   // 新周工作时长
	OriginalWorkShift          string                                                         `json:"original_work_shift,omitempty"`           // 原排班
	TargetWorkShift            string                                                         `json:"target_work_shift,omitempty"`             // 新排班
	OriginalCostCenterRate     []*CreateCoreHrJobChangeRespTransferInfoOriginalCostCenterRate `json:"original_cost_center_rate,omitempty"`     // 原成本中心分摊信息
	TargetCostCenterRate       []*CreateCoreHrJobChangeRespTransferInfoTargetCostCenterRate   `json:"target_cost_center_rate,omitempty"`       // 新成本中心分摊信息
}

// CreateCoreHrJobChangeRespTransferInfoOriginalCostCenterRate ...
type CreateCoreHrJobChangeRespTransferInfoOriginalCostCenterRate struct {
	CostCenterID string `json:"cost_center_id,omitempty"` // 支持的成本中心id
	Rate         int64  `json:"rate,omitempty"`           // 分摊比例
}

// CreateCoreHrJobChangeRespTransferInfoTargetCostCenterRate ...
type CreateCoreHrJobChangeRespTransferInfoTargetCostCenterRate struct {
	CostCenterID string `json:"cost_center_id,omitempty"` // 支持的成本中心id
	Rate         int64  `json:"rate,omitempty"`           // 分摊比例
}

// createCoreHrJobChangeResp ...
type createCoreHrJobChangeResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *CreateCoreHrJobChangeResp `json:"data,omitempty"`
}
