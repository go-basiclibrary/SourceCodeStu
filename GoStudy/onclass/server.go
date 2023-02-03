package main

import (
	"fmt"
	"net/http"
)

type Server interface {
	Routable
	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler Handler
	root    Filter
}

func (s *sdkHttpServer) Route(method, pattern string, handlerFunc func(c *Context)) {
	handlerBaseOnMap := NewHandlerBaseOnMap()
	handlerBaseOnMap.Route(method, pattern, handlerFunc)
}

func (s *sdkHttpServer) Start(address string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := NewContext(w, r)
		s.root(ctx)
	})
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string, filters ...FilterBuilder) Server {
	handler := NewHandlerBaseOnMap()

	var root Filter = func(c *Context) {
		handler.ServeHTTP(c)
	}

	for i := len(filters) - 1; i >= 0; i-- {
		root = filters[i](root)
	}

	return &sdkHttpServer{
		Name:    name,
		handler: handler,
		root:    root,
	}
}

type signUpReq struct {
}

type commonResponse struct {
	Data int
}

func SignUp(c *Context) {
	req := &signUpReq{}
	err := c.ReadJson(req)
	if err != nil {
		c.BadReqJson(err)
		return
	}

	resp := &commonResponse{
		Data: 123,
	}

	err = c.OkJson(resp)
	if err != nil {
		fmt.Errorf("写入响应失败:%v", err)
	}
}
