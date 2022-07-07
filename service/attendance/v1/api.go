// Package attendance code generated by oapi sdk gen
package larkattendance

import (
	"bytes"
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/core"
)

// 构建业务域服务实例
func NewService(config *core.Config) *AttendanceService {
	a := &AttendanceService{config: config}
	a.ApprovalInfo = &approvalInfo{service: a}
	a.File = &file{service: a}
	a.Group = &group{service: a}
	a.Shift = &shift{service: a}
	a.UserApproval = &userApproval{service: a}
	a.UserDailyShift = &userDailyShift{service: a}
	a.UserFlow = &userFlow{service: a}
	a.UserSetting = &userSetting{service: a}
	a.UserStatsData = &userStatsData{service: a}
	a.UserStatsField = &userStatsField{service: a}
	a.UserStatsView = &userStatsView{service: a}
	a.UserTask = &userTask{service: a}
	a.UserTaskRemedy = &userTaskRemedy{service: a}
	return a
}

// 业务域服务定义
type AttendanceService struct {
	config         *core.Config
	ApprovalInfo   *approvalInfo
	File           *file
	Group          *group
	Shift          *shift
	UserApproval   *userApproval
	UserDailyShift *userDailyShift
	UserFlow       *userFlow
	UserSetting    *userSetting
	UserStatsData  *userStatsData
	UserStatsField *userStatsField
	UserStatsView  *userStatsView
	UserTask       *userTask
	UserTaskRemedy *userTaskRemedy
}

// 资源服务定义
type approvalInfo struct {
	service *AttendanceService
}
type file struct {
	service *AttendanceService
}
type group struct {
	service *AttendanceService
}
type shift struct {
	service *AttendanceService
}
type userApproval struct {
	service *AttendanceService
}
type userDailyShift struct {
	service *AttendanceService
}
type userFlow struct {
	service *AttendanceService
}
type userSetting struct {
	service *AttendanceService
}
type userStatsData struct {
	service *AttendanceService
}
type userStatsField struct {
	service *AttendanceService
}
type userStatsView struct {
	service *AttendanceService
}
type userTask struct {
	service *AttendanceService
}
type userTaskRemedy struct {
	service *AttendanceService
}

// 资源服务方法定义
func (a *approvalInfo) Process(ctx context.Context, req *ProcessApprovalInfoReq, options ...core.RequestOptionFunc) (*ProcessApprovalInfoResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, a.service.config, http.MethodPost,
		"/open-apis/attendance/v1/approval_infos/process", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ProcessApprovalInfoResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) Download(ctx context.Context, req *DownloadFileReq, options ...core.RequestOptionFunc) (*DownloadFileResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodGet,
		"/open-apis/attendance/v1/files/:file_id/download", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DownloadFileResp{RawResponse: rawResp}
	// 如果是下载，则设置响应结果
	if rawResp.StatusCode == http.StatusOK {
		resp.File = bytes.NewBuffer(rawResp.RawBody)
		resp.FileName = core.FileNameByHeader(rawResp.Header)
		return resp, err
	}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) Upload(ctx context.Context, req *UploadFileReq, options ...core.RequestOptionFunc) (*UploadFileResp, error) {
	options = append(options, core.WithFileUpload())
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodPost,
		"/open-apis/attendance/v1/files/upload", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UploadFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (g *group) Create(ctx context.Context, req *CreateGroupReq, options ...core.RequestOptionFunc) (*CreateGroupResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, g.service.config, http.MethodPost,
		"/open-apis/attendance/v1/groups", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateGroupResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (g *group) Delete(ctx context.Context, req *DeleteGroupReq, options ...core.RequestOptionFunc) (*DeleteGroupResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, g.service.config, http.MethodDelete,
		"/open-apis/attendance/v1/groups/:group_id", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteGroupResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (g *group) Get(ctx context.Context, req *GetGroupReq, options ...core.RequestOptionFunc) (*GetGroupResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, g.service.config, http.MethodGet,
		"/open-apis/attendance/v1/groups/:group_id", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetGroupResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (g *group) List(ctx context.Context, req *ListGroupReq, options ...core.RequestOptionFunc) (*ListGroupResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, g.service.config, http.MethodGet,
		"/open-apis/attendance/v1/groups", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListGroupResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (g *group) ListByIterator(ctx context.Context, req *ListGroupReq, options ...core.RequestOptionFunc) (*ListGroupIterator, error) {
	return &ListGroupIterator{
		ctx:      ctx,
		req:      req,
		listFunc: g.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (g *group) Search(ctx context.Context, req *SearchGroupReq, options ...core.RequestOptionFunc) (*SearchGroupResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, g.service.config, http.MethodPost,
		"/open-apis/attendance/v1/groups/search", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SearchGroupResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *shift) Create(ctx context.Context, req *CreateShiftReq, options ...core.RequestOptionFunc) (*CreateShiftResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, s.service.config, http.MethodPost,
		"/open-apis/attendance/v1/shifts", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateShiftResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *shift) Delete(ctx context.Context, req *DeleteShiftReq, options ...core.RequestOptionFunc) (*DeleteShiftResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, s.service.config, http.MethodDelete,
		"/open-apis/attendance/v1/shifts/:shift_id", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteShiftResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *shift) Get(ctx context.Context, req *GetShiftReq, options ...core.RequestOptionFunc) (*GetShiftResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, s.service.config, http.MethodGet,
		"/open-apis/attendance/v1/shifts/:shift_id", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetShiftResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *shift) List(ctx context.Context, req *ListShiftReq, options ...core.RequestOptionFunc) (*ListShiftResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, s.service.config, http.MethodGet,
		"/open-apis/attendance/v1/shifts", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListShiftResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *shift) ListByIterator(ctx context.Context, req *ListShiftReq, options ...core.RequestOptionFunc) (*ListShiftIterator, error) {
	return &ListShiftIterator{
		ctx:      ctx,
		req:      req,
		listFunc: s.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (s *shift) Query(ctx context.Context, req *QueryShiftReq, options ...core.RequestOptionFunc) (*QueryShiftResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, s.service.config, http.MethodPost,
		"/open-apis/attendance/v1/shifts/query", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryShiftResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userApproval) Create(ctx context.Context, req *CreateUserApprovalReq, options ...core.RequestOptionFunc) (*CreateUserApprovalResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, u.service.config, http.MethodPost,
		"/open-apis/attendance/v1/user_approvals", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateUserApprovalResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userApproval) Query(ctx context.Context, req *QueryUserApprovalReq, options ...core.RequestOptionFunc) (*QueryUserApprovalResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, u.service.config, http.MethodPost,
		"/open-apis/attendance/v1/user_approvals/query", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUserApprovalResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userDailyShift) BatchCreate(ctx context.Context, req *BatchCreateUserDailyShiftReq, options ...core.RequestOptionFunc) (*BatchCreateUserDailyShiftResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, u.service.config, http.MethodPost,
		"/open-apis/attendance/v1/user_daily_shifts/batch_create", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchCreateUserDailyShiftResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userDailyShift) Query(ctx context.Context, req *QueryUserDailyShiftReq, options ...core.RequestOptionFunc) (*QueryUserDailyShiftResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, u.service.config, http.MethodPost,
		"/open-apis/attendance/v1/user_daily_shifts/query", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUserDailyShiftResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userFlow) BatchCreate(ctx context.Context, req *BatchCreateUserFlowReq, options ...core.RequestOptionFunc) (*BatchCreateUserFlowResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, u.service.config, http.MethodPost,
		"/open-apis/attendance/v1/user_flows/batch_create", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchCreateUserFlowResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userFlow) Get(ctx context.Context, req *GetUserFlowReq, options ...core.RequestOptionFunc) (*GetUserFlowResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, u.service.config, http.MethodGet,
		"/open-apis/attendance/v1/user_flows/:user_flow_id", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetUserFlowResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userFlow) Query(ctx context.Context, req *QueryUserFlowReq, options ...core.RequestOptionFunc) (*QueryUserFlowResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, u.service.config, http.MethodPost,
		"/open-apis/attendance/v1/user_flows/query", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUserFlowResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userSetting) Modify(ctx context.Context, req *ModifyUserSettingReq, options ...core.RequestOptionFunc) (*ModifyUserSettingResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, u.service.config, http.MethodPost,
		"/open-apis/attendance/v1/user_settings/modify", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ModifyUserSettingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userSetting) Query(ctx context.Context, req *QueryUserSettingReq, options ...core.RequestOptionFunc) (*QueryUserSettingResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, u.service.config, http.MethodGet,
		"/open-apis/attendance/v1/user_settings/query", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUserSettingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userStatsData) Query(ctx context.Context, req *QueryUserStatsDataReq, options ...core.RequestOptionFunc) (*QueryUserStatsDataResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, u.service.config, http.MethodPost,
		"/open-apis/attendance/v1/user_stats_datas/query", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUserStatsDataResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userStatsField) Query(ctx context.Context, req *QueryUserStatsFieldReq, options ...core.RequestOptionFunc) (*QueryUserStatsFieldResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, u.service.config, http.MethodPost,
		"/open-apis/attendance/v1/user_stats_fields/query", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUserStatsFieldResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userStatsView) Query(ctx context.Context, req *QueryUserStatsViewReq, options ...core.RequestOptionFunc) (*QueryUserStatsViewResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, u.service.config, http.MethodPost,
		"/open-apis/attendance/v1/user_stats_views/query", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUserStatsViewResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userStatsView) Update(ctx context.Context, req *UpdateUserStatsViewReq, options ...core.RequestOptionFunc) (*UpdateUserStatsViewResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, u.service.config, http.MethodPut,
		"/open-apis/attendance/v1/user_stats_views/:user_stats_view_id", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateUserStatsViewResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userTask) Query(ctx context.Context, req *QueryUserTaskReq, options ...core.RequestOptionFunc) (*QueryUserTaskResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, u.service.config, http.MethodPost,
		"/open-apis/attendance/v1/user_tasks/query", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUserTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userTaskRemedy) Create(ctx context.Context, req *CreateUserTaskRemedyReq, options ...core.RequestOptionFunc) (*CreateUserTaskRemedyResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, u.service.config, http.MethodPost,
		"/open-apis/attendance/v1/user_task_remedys", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateUserTaskRemedyResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userTaskRemedy) Query(ctx context.Context, req *QueryUserTaskRemedyReq, options ...core.RequestOptionFunc) (*QueryUserTaskRemedyResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, u.service.config, http.MethodPost,
		"/open-apis/attendance/v1/user_task_remedys/query", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUserTaskRemedyResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userTaskRemedy) QueryUserAllowedRemedys(ctx context.Context, req *QueryUserAllowedRemedysUserTaskRemedyReq, options ...core.RequestOptionFunc) (*QueryUserAllowedRemedysUserTaskRemedyResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, u.service.config, http.MethodPost,
		"/open-apis/attendance/v1/user_task_remedys/query_user_allowed_remedys", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUserAllowedRemedysUserTaskRemedyResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
