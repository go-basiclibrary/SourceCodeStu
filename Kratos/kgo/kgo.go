package main

import (
	"fmt"
	"unicode"
)

func main() {
	//err := trpc.Go(context.TODO(), 1*time.Second, func(ctx context.Context) {
	//	time.Sleep(5e9)
	//	_, ok := ctx.Deadline()
	//	if ok {
	//		fmt.Println("ctx cancel")
	//	}
	//	fmt.Println("i'm go")
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//select {}

	//value, found := syscall.Getenv("MySQL")
	//if found {
	//	fmt.Println(value)
	//}
	//fmt.Println(value)

	//env := trpc.ExpandEnv("${MySQL}")
	//fmt.Println(env)
	var s rune = 'a'
	//for _, v := range s {
	//	fmt.Println(string(v))
	//}

	upper := unicode.ToUpper(s)
	fmt.Println(string(upper))
}
