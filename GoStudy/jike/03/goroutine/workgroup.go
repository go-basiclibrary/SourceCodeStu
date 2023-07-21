package main

import (
	"context"
	"fmt"
	"net/http"
)

func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		// 加入生命周期控制,关闭机制
		<-stop
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

func main() {
	done := make(chan error, 2)
	stop := make(chan struct{})
	go func() {
		done <- serve(":8001", http.NewServeMux(), stop)
	}()
	go func() {
		done <- serveApp(stop)
	}()

	var stopped bool
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Printf("error: %v\n", err)
		}
		if !stopped {
			stopped = true
			// 关闭资源
			close(stop)
		}
	}
}

func serveApp(stop chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Qcon!!!")
	})
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		stop <- struct{}{}
		return err
	}
	return nil
}

func serverDebug(stop chan struct{}) error {
	err := http.ListenAndServe(":8001", http.DefaultServeMux)
	if err != nil {
		stop <- struct{}{}
		return err
	}

	return nil
}
