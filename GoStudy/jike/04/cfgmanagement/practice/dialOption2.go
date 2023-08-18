package main

import "fmt"

type obj struct {
	allThings string
}

type optionO func(o *obj) optionO

func (o *obj) Option(opts ...optionO) (previous optionO) {
	for _, opt := range opts {
		previous = opt(o)
	}
	return
}

func ChangeAllTings(something string) optionO {
	return func(o *obj) optionO {
		prev := o.allThings
		o.allThings = something
		return ChangeAllTings(prev)
	}
}

func DoSomethingChange(o *obj, something string) {
	prev := o.Option(ChangeAllTings(something))
	defer prev(o)
	fmt.Printf("i'm change,%v\n", o.allThings)
}

func main() {
	o := &obj{allThings: "小丑"}
	fmt.Printf("i'm init,%v\n", o.allThings)

	// 中间需要做一些事情,对于某种特定状态进行变更
	DoSomethingChange(o, "我是万物之主!")
	fmt.Printf("i'm callback,%v\n", o.allThings)
}
