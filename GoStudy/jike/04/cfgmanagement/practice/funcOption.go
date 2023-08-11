package main

import "fmt"

type Foo struct {
	verify int
}

type option func(*Foo) option

func (f *Foo) Option(option ...option) {
	for _, o := range option {
		o(f)
	}
}

func Verbosity(v int) option {
	return func(f *Foo) option {
		prev := f.verify
		f.verify = v
		return Verbosity(prev)
	}
}

func DoSomething(f *Foo, v int) {
	// 修改后做一些事情
	op := Verbosity(v)
	fmt.Println("变更前,什么情况,", f.verify)
	o := op(f)
	fmt.Println("变更后,做一些事情,", f.verify)
	defer o(f)
}

func main() {
	f := &Foo{verify: 100}
	fmt.Println("初始化,长什么样子,", f.verify)
	DoSomething(f, 101)
	fmt.Println("最终,我变成了什么样子,", f.verify)
}
