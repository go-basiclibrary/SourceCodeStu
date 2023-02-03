package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 过滤包含中文的字符串
	m, err := regexp.MatchString("^[^\u4e00-\u9fa5]*$", "123王abc")
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
}
