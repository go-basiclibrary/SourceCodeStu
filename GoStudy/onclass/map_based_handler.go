package main

import (
	"net/http"
)

type Routable interface {
	Route(method, pattern string, handlerFunc func(c *Context))
}

type Handler interface {
	ServeHTTP(c *Context)
	Routable
}

type HandlerBasedOnMap struct {
	// method + url
	handlers map[string]func(ctx *Context)
}

// Route 将server端动作进行委托,否则需要暴露handlers
func (h *HandlerBasedOnMap) Route(method, pattern string, handlerFunc func(c *Context)) {
	key := h.key(method, pattern)
	h.handlers[key] = handlerFunc
}

// ServeHTTP
func (h *HandlerBasedOnMap) ServeHTTP(c *Context) {
	if handler, ok := h.handlers[h.key(c.R.Method, c.R.URL.Path)]; ok {
		handler(c)
	} else {
		// register
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("Not Found"))
	}
}

// key
func (h HandlerBasedOnMap) key(method, path string) string {
	return method + "#" + path
}

// 确保 HandlerBasedOnMap 一定实现了 Handler  或者左边的类型是否等于右边的类型
var _ Handler = &HandlerBasedOnMap{}

func NewHandlerBaseOnMap() Handler {
	return &HandlerBasedOnMap{handlers: make(map[string]func(c *Context))}
}
