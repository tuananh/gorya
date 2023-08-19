package handler

import (
	"context"
	"encoding/json"
	"github.com/nduyphuong/gorya/internal/store"
	svcv1alpha1 "github.com/nduyphuong/gorya/pkg/api/service/v1alpha1"
	"net/http"
)

func GetScheduleV1alpha1(ctx context.Context, store store.Interface) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("schedule")
		schedule, err := store.GetSchedule(name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		resp := svcv1alpha1.GetScheduleResponse{
			ScheduleModel: schedule,
		}
		b, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)
	}
}
