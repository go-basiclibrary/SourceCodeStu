package main

import (
	_ "net/http/pprof"
)

func main() {
	//mux := http.DefaultServeMux
	//fmt.Println(*mux)
	//err := http.ListenAndServe("127.0.0.1:8081", mux)
	//if err != nil {
	//	panic(err)
	//}

	//var ch = make(chan int, 20)
	//go func() {
	//	ch <- 55
	//	ch <- 2
	//	close(ch)
	//}()
	//time.Sleep(1e9)
	//for {
	//	time.Sleep(1e9)
	//	select {
	//	case v := <-ch:
	//		fmt.Println(v)
	//	}
	//}
}
