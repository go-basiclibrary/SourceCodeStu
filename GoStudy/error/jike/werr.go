package main

import (
	"git.tencent.com/trpc-go/trpc-go/log"
	"time"
)

func main() {
	//err := errors.New("abc")
	//err = fmt.Errorf("%s: %w", "/jike/werr", err)
	//err = errors.Unwrap(err)
	//fmt.Println(err)

	now := time.Now()
	time.Sleep(1 * time.Second)
	since := time.Since(now)
	log.Infof("%s", since)
}
