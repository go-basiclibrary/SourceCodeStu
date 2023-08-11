package main

import "fmt"

// https://gist.github.com/yulvil/5c4d4903bf0b0a332003

type Foo struct {
	verbosity int
}

func (f *Foo) Option(opts ...option) (previous option) {
	for _, opt := range opts {
		previous = opt(f)
	}
	return previous
}

type option func(f *Foo) option

func Verbosity(v int) option {
	return func(f *Foo) option {
		prev := f.verbosity
		f.verbosity = v
		return Verbosity(prev)
	}
}

func DoSomethingVerbosely(f *Foo, verbosity int) {
	prev := f.Option(Verbosity(verbosity))
	defer f.Option(prev)

	// verbosity一定变了
	fmt.Printf("second verbosity is %d\n", f.verbosity)

	// defer后verbosity回归
}

func main() {
	f := &Foo{
		verbosity: 0,
	}
	fmt.Printf("first verbosity is %d\n", f.verbosity)
	DoSomethingVerbosely(f, 1)
	fmt.Printf("third verbosity is %d\n", f.verbosity)
}
