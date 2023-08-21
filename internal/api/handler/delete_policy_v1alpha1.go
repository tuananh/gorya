package handler

import (
	"context"
	"github.com/nduyphuong/gorya/internal/store"
	"net/http"
)

func DeletePolicyV1alpha1(ctx context.Context, store store.Interface) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("policy")
		if isEmpty(name) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := store.DeletePolicy(name); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
