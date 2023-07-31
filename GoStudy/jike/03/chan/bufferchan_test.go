package main

import (
	"testing"
)

func BenchmarkWithNoBuffer(b *testing.B) {
	benchmarkWithBuffer(b, 0)
}

func BenchmarkWithBufferSizeOf1(b *testing.B) {
	benchmarkWithBuffer(b, 1)
}

func BenchmarkWithBufferEqualsToNumberOfWorker(b *testing.B) {
	benchmarkWithBuffer(b, 5)
}

func BenchmarkWithBufferSizeExceedsNumberOfWorker(b *testing.B) {
	benchmarkWithBuffer(b, 25)
}

func benchmarkWithBuffer(b *testing.B, bufferSize int) {
	stopChan := make(chan struct{})

	iChan := make(chan int, bufferSize)
	// 制造百万数据
	go func() {
		for i := 0; i <= 1000000; i++ {
			iChan <- i
		}
		close(iChan)
		stopChan <- struct{}{}
	}()

	for i := 0; i < 5; i++ {
		go fn(iChan)
	}

	<-stopChan
}

func fn(iChan chan int) {
	for {
		switch <-iChan {
		}
	}
}
