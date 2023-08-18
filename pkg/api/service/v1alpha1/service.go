package v1alpha1

import (
	"context"
	connect_go "github.com/bufbuild/connect-go"
	"github.com/nduyphuong/gorya/internal/store"
	"net/http"
	"time"
)

type GoryaServiceHandler interface {
	GetTimeZone() http.Handler
	GetVersionInfo() http.Handler
	AddSchedule(ctx context.Context, store store.Interface) http.Handler
}

const (
	GoryaTaskChangeStageProcedure = "/tasks/change_state"
	GoryaTaskGetScheduleProcedure = "/tasks/schedule"
	GoryaGetTimeZoneProcedure     = "/api/v1alpha1/time_zones"
	GoryaAddScheduleProcedure     = "/api/v1alpha1/add_schedule"
	GoryaGetScheduleProcedure     = "/api/v1alpha1/get_schedule"
	GoryaListScheduleProcedure    = "/api/v1alpha1/list_schedules"
	GoryaDeleteScheduleProcedure  = "/api/v1alpha1/del_schedule"
	GoryaAddPolicyProcedure       = "/api/v1alpha1/add_policy"
	GoryaGetPolicyProcedure       = "/api/v1alpha1/get_policy"
	GoryaListPolicyProcedure      = "/api/v1alpha1/list_policies"
	GoryaDeletePolicyProcedure    = "/api/v1alpha1/del_policy"
	GoryaGetVersionInfo           = "/api/v1alpha1/version_info"
)

// NewGoryaServiceHandler builds an HTTP handler from the service implementation. It returns the
//
//	path on which to mount the handler and the handler itself.
func NewGoryaServiceHandler(ctx context.Context, store store.Interface, svc GoryaServiceHandler,
	opts ...connect_go.HandlerOption) (string,
	http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(GoryaGetTimeZoneProcedure, svc.GetTimeZone())
	mux.Handle(GoryaGetVersionInfo, svc.GetVersionInfo())
	mux.Handle(GoryaAddScheduleProcedure, svc.AddSchedule(ctx, store))
	return "/", mux
}

type VersionInfo struct {
	Version      string    ` json:"version,omitempty"`
	GitCommit    string    `json:"git_commit,omitempty"`
	GitTreeDirty bool      `json:"git_tree_dirty,omitempty"`
	BuildTime    time.Time `json:"build_time,omitempty"`
	GoVersion    string    `json:"go_version,omitempty"`
	Compiler     string    `json:"compiler,omitempty"`
	Platform     string    `json:"platform,omitempty"`
}

type OkResponse struct {
	Message string `json:"message"`
}

type GetVersionInfoResponse struct {
	VersionInfo *VersionInfo `json:"version_info,omitempty"`
}

type GetTimeZoneResponse struct {
	TimeZones []string `json:"Timezones,omitempty"`
}

type AddScheduleRequest struct {
	Name        string  `json:"name"`
	DisplayName string  `json:"displayname,omitempty"`
	Dtype       string  `json:"dtype"`
	Corder      bool    `json:"corder"`
	Shape       []int   `json:"shape"`
	NdArray     [][]int `json:"__ndarray__"`
	TimeZone    string  `json:"timezone"`
}

type AddPolicyRequest struct {
	Name        string              `json:"name"`
	DisplayName string              `json:"displayname,omitempty"`
	Tags        []map[string]string `json:"tags"`
	//compatible with front-end logic
	Accounts     []string `json:"projects"`
	ScheduleName string   `json:"schedulename"`
}
