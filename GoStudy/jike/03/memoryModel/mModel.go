package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var num int

func main() {
	for i := 0; i < 100; i++ {
		var a atomic.Value
		go func() {
			a.Store(1)
		}()
		go func() {
			if a.Load() == 0 {
				num++
				fmt.Printf("the result is %d,i %d\n", a, num)
			}
		}()
	}
	time.Sleep(1e9)
}
