package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var i int
	done := make(chan bool, 1)
	var m sync.Mutex
	// g1
	go func() {
		for {
			i++
			select {
			case <-done:
				fmt.Printf("g1 is run time %d\n", i)
				return
			default:
				m.Lock()
				time.Sleep(100 * time.Millisecond)
				m.Unlock()
			}
		}
	}()

	// g2
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		m.Lock()
		m.Unlock()
	}
	done <- true
	time.Sleep(1e9)
}
