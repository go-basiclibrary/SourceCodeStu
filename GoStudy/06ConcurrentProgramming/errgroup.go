package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	errgroup, _ := errgroup.WithContext(context.TODO())
	var i *int
	i = new(int)
	errgroup.Go(func() error {
		*i = 1
		return nil
	})
	errgroup.Wait()
	fmt.Println(*i)

	//group := singleflight.Group{}
	//v, err, shared := group.Do("1", func() (interface{}, error) {
	//	return 1, nil
	//})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(v, shared)
}
