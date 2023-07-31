package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup
var Counter int = 0

// 全局变量并发造成data race
func main() {
	for routine := 0; routine < 2; routine++ {
		wait.Add(1)
		go Routine(routine)
	}
	wait.Wait()
	fmt.Println(Counter)
}

func Routine(id int) {
	for count := 0; count < 2; count++ {
		value := Counter
		//time.Sleep(1 * time.Nanosecond)

		// 该代码汇编方式
		// MOVQ Copy出来
		// INCQ 自增
		// MOVQ 移动新的value回去
		value++
		Counter = value
	}

	wait.Done()
}
