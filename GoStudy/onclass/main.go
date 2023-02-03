package main

import (
	"net/http"
)

func main() {
	server := NewHttpServer("test-server")
	//server.Route("/", home)
	server.Route(http.MethodGet, "/user/signup", SignUp)
	err := server.Start(":8080")
	if err != nil {
		panic(err)
	}
}
