package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a string
	var b int
	var c struct{}
	var c2 *struct{}
	var d *int
	fmt.Printf("%p\n", &a)
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(c))
	fmt.Printf("%p\n", &c)
	fmt.Printf("%p\n", c2)
	fmt.Printf("%p\n", d)
}
