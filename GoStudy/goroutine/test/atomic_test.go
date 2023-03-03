package test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

type Config struct {
	a []int
}

func (c *Config) T() {
}

func BenchmarkAtomic(b *testing.B) {
	var v atomic.Value
	v.Store(&Config{})

	go func() {
		i := 0
		for {
			i++
			cfg := &Config{[]int{i, i + 1, i + 2, i + 3, i + 4, i + 5}}
			v.Store(cfg)
		}
	}()

	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				cfg := v.Load().(*Config)
				cfg.T()
				fmt.Printf("%v\n", cfg)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkMutex(b *testing.B) {
	var l sync.Mutex
	var cfg *Config

	go func() {
		i := 0
		for {
			i++
			l.Lock()
			cfg = &Config{[]int{i, i + 1, i + 2, i + 3, i + 4, i + 5}}
			l.Unlock()
		}
	}()

	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				l.Lock()
				fmt.Printf("%v\n", cfg)
				l.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
