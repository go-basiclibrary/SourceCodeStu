package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 并发访问 no single machine world type is no safe
// 使用atomic 进行更改代码
func main() {
	var v atomic.Value
	v.Store(&config{})
	//cfg := &config{}

	go func() {
		i := 0
		for {
			i++
			// slice is not concurrency safe
			// cfg.a = []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}

			// use atomic
			cfg := &config{a: []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}}
			v.Store(cfg)
		}
	}()

	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				//fmt.Printf("%v\n", cfg)
				cfg := v.Load().(*config)
				fmt.Printf("%v\n", cfg)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

type config struct {
	a []int
}
