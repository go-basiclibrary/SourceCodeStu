package main

import (
	"fmt"
	"sync"
	"time"
)

var Wait sync.WaitGroup
var Counter int = 0

func main() {
	for routine := 0; routine < 2; routine++ {
		Wait.Add(1)
		go Routine(routine)
	}
	Wait.Wait()
	fmt.Printf("Final Counter: %d\n", Counter)
}

func Routine(id int) {
	for count := 0; count < 2; count++ {
		// 改行实际看起来只有一行代码
		// 但是转变成汇编以后,是三行
		// Copy -> +1 -> 将新数据移动回Counter
		Counter = Counter + 1 // Read Write 被多个进程同时进行
		time.Sleep(1 * time.Nanosecond)
	}

	Wait.Done()
}
