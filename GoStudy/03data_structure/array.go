package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic")
		}
	}()
	// 运行时panic
	var arr [3]int = [3]int{1, 2, 3}
	i := 1
	i++
	i++
	i++
	fmt.Println(arr[i])
}
