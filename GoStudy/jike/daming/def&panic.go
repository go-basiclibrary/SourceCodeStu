package main

import "fmt"

func main() {
	t()()
}

func t() func() {
	x := 1
	f := func() {
		fmt.Println(x)
	}
	x = 2
	return f
}
