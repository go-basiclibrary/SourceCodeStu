package main

import "fmt"

func main() {
	// 类型转换
	var i interface{} = 10
	fmt.Println(i.(string))
}
