// Code generated by protoc-gen-go-http. DO NOT EDIT.

package v1

import (
	context "context"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	mux "github.com/gorilla/mux"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(http.Request)
var _ = new(context.Context)
var _ = binding.MapProto
var _ = mux.NewRouter

const _ = http1.SupportPackageIsVersion1

type UserHandler interface {
	GetMyMessages(context.Context, *GetMyMessagesRequest) (*GetMyMessagesReply, error)
}

func NewUserHandler(srv UserHandler, opts ...http1.HandleOption) http.Handler {
	h := http1.DefaultHandleOptions()
	for _, o := range opts {
		o(&h)
	}
	r := mux.NewRouter()

	r.HandleFunc("/v1/user/get/message/{count}", func(w http.ResponseWriter, r *http.Request) {
		var in GetMyMessagesRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := binding.MapProto(&in, mux.Vars(r)); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMyMessages(ctx, req.(*GetMyMessagesRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*GetMyMessagesReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	return r
}
