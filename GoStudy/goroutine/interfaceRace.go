package main

import "fmt"

// 看起来下面的代码,并没有产生什么原子性
// 实则,当我们了解interface底层类型以后, *type,*data形成interface
// 当并发进行转换时,可能另一个协程在转换*type还没来得及转换*data的时候,就会造成并发 data race
func main() {
	var ben = &Ben{10, "Ben"}
	var jerry = &Jerry{"Jerry"}
	var maker IceCreamMaker = ben

	var loop0, loop1 func()

	loop0 = func() {
		// 在maker进行转换的时候,*type *data在汇编中是两个原子操作
		maker = ben
		go loop1()
	}
	loop1 = func() {
		maker = jerry
		go loop0()
	}

	go loop0()

	for {
		maker.Hello()
	}
}

type IceCreamMaker interface {
	Hello()
}

type Ben struct {
	id   int
	name string
}

func (b *Ben) Hello() {
	fmt.Printf("Ben says,\"Hello my name is %s\"\n", b.name)
}

type Jerry struct {
	name string
}

func (j *Jerry) Hello() {
	fmt.Printf("Jerry says,\"Hello my name is %s\"\n", j.name)
}
