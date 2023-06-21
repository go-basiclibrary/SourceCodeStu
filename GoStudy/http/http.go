package main

import (
	"fmt"
	"net/url"
)

func main() {
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//defer fmt.Println(3)
	//panic(123)
	//
	//// 断言
	//var a interface{}
	//i := a.(int)
	//fmt.Println(i)

	//var slice []int
	//fmt.Println(slice == nil)
	//
	//escape, err := url.QueryUnescape(URL)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(escape)

	//sum := md5.Sum([]byte("xx+20302+wcg+王者农药"))
	//md5String := hex.EncodeToString(sum[:])
	//marshal, _ := proto.Marshal(md5String)
	//fmt.Println(string(marshal))

	s := fmt.Sprintf(URL, url.QueryEscape("https://test-bill-mdms.woa.com/bill-report/todolist/114?onlycontent=1"))
	fmt.Println(s)
}

var URL = "https://test-mdms.woa.com/index.html#/bill/billing-report?billadmin=%s"
