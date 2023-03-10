package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var i int32
	go func() {
		time.Sleep(0.5e9)
		res := atomic.AddInt32(&i, -1)
		fmt.Println(res)
		if res != 0 {
			fmt.Println("我不是0")
		}
	}()

	i = 1
	time.Sleep(3e9)
}
