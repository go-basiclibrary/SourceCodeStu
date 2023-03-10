package main

import "fmt"

func main() {
	//c1 := make(chan interface{})
	//c2 := make(chan interface{})
	//go func() {
	//	c1 <- 1
	//}()
	//go func() {
	//	c2 <- 1
	//}()
	//
	//select {
	//case v := <-c1:
	//	fmt.Printf("c1 is %v\n", v)
	//case v := <-c2:
	//	fmt.Printf("c2 is %v\n", v)
	//}

	// 直接阻塞 select {}
	//fmt.Println("start...")
	//select {}

	// 单一 Channel
	//var c = make(chan interface{})
	//go func() {
	//	c <- 1
	//}()
	//select {  // 非阻塞收发
	//case v, ok := <-c: // 当case的Channel是空指针的时候并且没有default,则直接panic
	//	fmt.Println(v, ok)
	//default:
	//	fmt.Println(5)
	//}

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, GopherCon SG")
	//})
	//go func() {
	//	if err := http.ListenAndServe(":8080", nil); err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	//
	//select {}

	// 缓冲chan
	//var c = make(chan int, 100)
	//go func() {
	//	c <- 1
	//}()
	//time.Sleep(1e9)
	//select {
	//case v := <-c:
	//	fmt.Println(v)
	//default:
	//	fmt.Println("default")
	//}

	// 单独一个select
	//select {}

	// c 为nil 时,该协程永远不会被唤醒
	var c chan int
	//c = make(chan int)
	go func() {
		c <- 1
	}()
	select {
	case <-c:
	default:
		fmt.Println("")
	}
}
