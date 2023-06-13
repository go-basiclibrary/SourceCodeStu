package main

import (
	"fmt"
	"time"
)

func main() {
	//deferFunc()
	//deferZz()
	//deferZx()

	//fmt.Println(*deferTT())

	//deferFuncTwice()

	deferFuncWithFor()
}

func deferFuncWithFor() {
	for i := 0; i < 3; i++ {
		procLoop()
		time.Sleep(1e9)
	}
}

func procLoop() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	proc()
}

func proc() {
	panic("proc")
}

func deferFuncTwice() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	testFunc()
}

func testFunc() {
	panic("test")
}

func deferTT() *int {
	t := 0
	defer func(i *int) { *i++ }(&t)
	return &t
}

// defer 执行顺序
func deferZx() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

type N struct {
	a int
}

// 预计算指针是否会有效果
// 是会影响值的
func deferZz() {
	var s = &N{a: 100}
	defer fmt.Println(s)
	s.a = 101
}

// defer 会进行预计算
func deferFunc() {
	startAt := time.Now()
	//defer fmt.Println(time.Since(startAt)) // 预期5s但是这里只会得到0s

	// 改进方法
	defer func() {
		fmt.Println(time.Since(startAt))
	}()
	time.Sleep(5e9)
}
