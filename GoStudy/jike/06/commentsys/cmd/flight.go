package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"jike/06/commentsys/api"
	"time"
)

func main() {
	// 模拟一百个人同时访问
	for i := 0; i < 100; i++ {
		go req("缓存穿透请求")
	}
	time.Sleep(2e9)
}

var sf singleflight.Group

func req(req string) {
	_, err, _ := sf.Do(req, func() (interface{}, error) {
		// 模拟存储
		res, err := reqDB(req)
		if err != nil {
			return nil, err
		}
		// 投递kfk进行回源
		err = deliverKfk(res)
		sf.Forget(req)
		return res, nil
	})
	if err != nil {
		panic(err)
	}
	//fmt.Println(v, err, shared)
}

func deliverKfk(res *api.Person) error {
	fmt.Println("投递kfk进行回源...")
	return nil
}

func reqDB(r string) (*api.Person, error) {
	fmt.Println("谁进来了")
	// 模拟db
	return &api.Person{
		Id:  1,
		Msg: &api.Msg{Msg: r},
	}, nil
}
