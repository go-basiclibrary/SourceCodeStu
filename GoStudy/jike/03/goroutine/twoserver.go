package main

// 启动两个server
func main() {
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintln(w, "Hello, QCon!")
	//})
	//go http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux)
	//http.ListenAndServe("0.0.0.0:8080", mux)

	// 将并发性交给调用者决定
	//go serverDebug()
	//go serverApp()

	select {}
}

// 优化启动两个server
//func serverApp() {
//	mux := http.NewServeMux()
//	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintln(w, "Hello, QCon!")
//	})
//	err := http.ListenAndServe("0.0.0.0:8080", mux)
//	if err != nil {
//		// 中断fatal 会造成defer等内置收尾函数不被执行
//		log.Fatal(err)
//	}
//}
//
//// debug server
//func serverDebug() {
//	// 无法感知 关闭状态等
//	err := http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
