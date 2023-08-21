package handler

import (
	"encoding/json"
	svcv1alpha1 "github.com/nduyphuong/gorya/pkg/api/service/v1alpha1"
	"github.com/nduyphuong/gorya/pkg/timezone"
	"net/http"
)

func GetTimeZoneV1Alpha1() http.HandlerFunc {
	resp := &svcv1alpha1.GetTimeZoneResponse{
		TimeZones: timezone.List(),
	}
	return func(w http.ResponseWriter, req *http.Request) {
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(b)
	}
}
