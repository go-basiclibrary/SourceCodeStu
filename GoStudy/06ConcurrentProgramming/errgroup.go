package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
)

func main() {
	//errgroup, ctx := errgroup.WithContext(context.TODO())
	//errgroup.Go(func() error {
	//	fmt.Println("123")
	//	return nil
	//})

	group := singleflight.Group{}
	v, err, shared := group.Do("1", func() (interface{}, error) {
		return 1, nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(v, shared)
}
