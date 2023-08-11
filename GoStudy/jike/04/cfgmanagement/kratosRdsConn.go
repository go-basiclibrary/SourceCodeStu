package main

import (
	"fmt"
	"github.com/go-kratos/kratos/pkg/cache/redis"
	"time"
)

func main() {
	c, _ := redis.Dial("tcp", "127.0.0.1:3389")
	redis.DialDatabase(0)
	redis.DialPassword("hello")
	redis.DialReadTimeout(10 * time.Second)

	fmt.Println(c)
}
