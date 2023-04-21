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

	//var i *int
	//*i = 10

	//ints := make(chan int)
	//go func() {
	//	fmt.Println("start ...")
	//	time.Sleep(5e9)
	//	i = 1
	//	go func() {
	//		for {
	//			fmt.Printf("select * from xx\n")
	//			select {
	//			case v := <-ints:
	//
	//			}
	//		}
	//	}()
	//	ints <- 1
	//}()
}

var i int32

type Stu struct {
	abc string
}
