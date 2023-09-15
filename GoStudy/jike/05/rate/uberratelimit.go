package main

import (
	"fmt"
	"go.uber.org/ratelimit"
	"time"
)

func main() {
	uberRateLimit(100)
}

// uber rate limit
func uberRateLimit(times int) {
	rl := ratelimit.New(times, ratelimit.WithoutSlack)

	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := rl.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}
