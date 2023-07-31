package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// parent ctx
	ctx := context.TODO()
	go func(ctx context.Context) {
		switch struct{}{} {
		case <-ctx.Done():
			fmt.Println("parent ctx is done")
		}
	}(ctx)

	// with time out child
	ctx, cancelFunc := context.WithCancel(ctx)
	go func(ctx context.Context) {
		switch struct{}{} {
		case <-ctx.Done():
			fmt.Println("child1 ctx is done")
		}
	}(ctx)

	ctx, _ = context.WithCancel(ctx)

	go func(ctx context.Context) {
		switch struct{}{} {
		case <-ctx.Done():
			fmt.Println("child2 ctx is done")
		}
	}(ctx)

	cancelFunc()

	time.Sleep(1 * time.Minute)
}
