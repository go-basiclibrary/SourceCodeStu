package main

import "net/http"

func main() {

}

type Server interface {
	Route(pattern string, handlerFunc http.HandlerFunc)
}

type sdkHttpServer struct {
	Name string
}
