package handler

import (
	"context"
	"encoding/json"
	"github.com/nduyphuong/gorya/internal/store"
	svcv1alpha1 "github.com/nduyphuong/gorya/pkg/api/service/v1alpha1"
	"gorm.io/gorm"
	"net/http"
)

func isEmpty(s string) bool {
	return s == ""
}
func GetScheduleV1alpha1(ctx context.Context, store store.Interface) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("schedule")
		if isEmpty(name) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		schedule, err := store.GetSchedule(name)
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusNotFound)
			return
		}
		resp := svcv1alpha1.GetScheduleResponse{
			ScheduleModel: *schedule,
		}
		b, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)
	}
}
