package main

import (
	"fmt"
	"sync"
	"time"
)

// mutex 互斥锁使用
func main() {
	done := make(chan bool, 1)
	var mu sync.Mutex

	var g1Count, g2Count int

	// g1
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				mu.Lock()
				g1Count++
				time.Sleep(100 * time.Millisecond)
				mu.Unlock()
			}
		}
	}()

	// g2
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		mu.Lock()
		g2Count++
		mu.Unlock()
	}

	done <- true
	fmt.Println(g1Count, "--", g2Count)
}

// g1->g2 30->10
