package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var num atomic.Value

func main() {
	// atomic改造,去除data race
	num.Store(0)

	var wg sync.WaitGroup
	// data race 情况
	for routine := 0; routine < 2; routine++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 2; i++ {
				val := num.Load()
				num.Store(val.(int) + 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Printf("Final Counter: %d\n", num.Load().(int))
}
