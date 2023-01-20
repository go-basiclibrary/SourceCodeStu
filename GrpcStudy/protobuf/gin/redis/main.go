package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", "121.37.232.8", 6379),
	})

	fmt.Println(*client)

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err.Error())
	}

	res, err := client.Set(context.Background(), "key", "key", 5*time.Minute).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
