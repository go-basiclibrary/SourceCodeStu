package main

import (
	"fmt"
	"time"
)

func main() {
	queue := func(work func(int)) {
		work(0)
	}

	workerIDs := make(chan int, 2)

	for i := 0; i < 2; i++ {
		workerIDs <- i
	}

	if true {
		queue = func(work func(int)) {
			go func() {
				worker := <-workerIDs
				work(worker)
				workerIDs <- worker
			}()
		}
	}

	queue(func(i int) {
		fmt.Println(i)
	})
	time.Sleep(1e9)
}
