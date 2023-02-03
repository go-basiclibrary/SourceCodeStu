package main

import (
	"fmt"
	"time"
)

type handlerFunc Filter

type FilterBuilder func(next Filter) Filter

type Filter func(c *Context)

// 确保 FilterBuilder 和 MetricsFilterBuilder 是同一个类型
var _ FilterBuilder = MetricsFilterBuilder

func MetricsFilterBuilder(next Filter) Filter {
	return func(c *Context) {
		start := time.Now().Nanosecond()
		next(c)
		end := time.Now().Nanosecond()
		fmt.Printf("用了 %d 纳秒\n", end-start)
	}
}
