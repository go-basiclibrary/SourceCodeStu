package main

import (
	"context"
	"fmt"
	"time"
)

type handlerFunc Filter

type FilterBuilder func(next Filter) Filter

type Filter func(c *context.Context)

// 确保 FilterBuilder 和 MetricsFilterBuilder 是同一个类型
var _ FilterBuilder = MetricsFilterBuilder

func MetricsFilterBuilder(next Filter) Filter {
	return func(c *context.Context) {
		start := time.Now().Nanosecond()
		next(c)
		end := time.Now().Nanosecond()
		fmt.Printf("用了 %d 纳秒\n", end-start)
	}
}

func main() {

}

func RegisterFilter(filters ...FilterBuilder) {

}
