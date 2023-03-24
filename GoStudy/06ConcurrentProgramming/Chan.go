package main

func main() {
	//var c chan int
	//
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	panic(err)
	//}
	//
	//<-c

	//chInt := make(chan int)
	//
	//go func() {
	//	for i := 0; i < 3; i++ {
	//		chInt <- i
	//	}
	//	close(chInt)
	//}()
	//for i := 0; i < 3; i++ {
	//	time.Sleep(0.5e9)
	//	go func(i int) {
	//		v := <-chInt
	//		fmt.Printf("%d:%d\n", i, v)
	//	}(i)
	//}
	//
	//time.Sleep(1e9)

	var i *int
	*i = 10
}

type Stu struct {
	abc string
}
