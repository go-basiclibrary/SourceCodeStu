package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// 永远不要在不知道什么时候会停止的情况下开启协程

type result struct {
	record string
	err    error
}

func main() {
	ch := make(chan result)
	ctx := context.TODO()
	go func() {
		record, err := search("")
		ch <- result{record, err}
	}()

	select {
	case <-ctx.Done(): // 监听ctx是否cancel
		fmt.Println(errors.New("search canceled"))
	case resultChan := <-ch:
		if resultChan.err != nil {
			return
		}
		fmt.Println("Received:", resultChan.record)
	}
}

// search函数,模拟实现,用于模拟长时间运行操作,如db or rpc
func search(term string) (string, error) {
	time.Sleep(200 * time.Millisecond)
	return "some value", nil
}

func process(term string) error {
	record, err := search(term)
	if err != nil {
		return err
	}

	fmt.Println("Received:", record)
	return nil
}
