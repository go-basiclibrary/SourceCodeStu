package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//var wg sync.WaitGroup
	//wg = sync.WaitGroup{}
	//transmit(wg)
	//wgCopy := wg
	//wg.Add(1)
	//fmt.Println(wg, wgCopy)

	//wg := sync.WaitGroup{}

	// once
	//once := sync.Once{}
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		once.Do(func() {
	//			fmt.Println(i)
	//		})
	//	}
	//}()
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		once.Do(func() {
	//			fmt.Println(i)
	//		})
	//	}
	//}()
	//time.Sleep(1e9)

	// Cond
	//c := sync.NewCond(&sync.Mutex{})
	//for i := 0; i < 10; i++ {
	//	go listen(c)
	//}
	//time.Sleep(1e9)
	//go broadcast(c)
	//
	//ch := make(chan os.Signal, 1)
	//signal.Notify(ch, os.Interrupt)
	//<-ch

	// copy wg
	//var wg sync.WaitGroup
	//wg.Add(1)
	//wg.Add(1)
	//wg.Done()
	//wg.Done()
	//wg.Wait()
	//var wg sync.WaitGroup
	//wg.Add(-1)

	var once sync.Once
	for i := 0; i < 3; i++ {
		go func() {
			once.Do(func() {
				fmt.Println(i)
			})
		}()
	}

	time.Sleep(1e9)
}

func listen(c *sync.Cond) {
	c.L.Lock()
	for atomic.LoadInt64(&status) != 1 {
		c.Wait() // 等待满足特定条件
	}
	fmt.Println("listen")
	c.L.Unlock()
}

var status int64

func broadcast(c *sync.Cond) {
	c.L.Lock()
	atomic.StoreInt64(&status, 1) // 更新条件,唤醒陷入等待的协程
	c.Signal()
	c.L.Unlock()
}

func transmit(wg sync.WaitGroup) {
	fmt.Println(wg)
}
