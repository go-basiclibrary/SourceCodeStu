package main

import (
	"fmt"
)

// 反射
func main() {
	//var d = &DoG{}
	// 通过反射必须要这么才能判断是否实现了该接口
	//bo := reflect.TypeOf(d).Implements(reflect.TypeOf((*DucK)(nil)).Elem())
	//if bo {
	//	fmt.Println("我实现了 DucK")
	//}

	// 第一法则  从interface{}到反射对象
	//var i interface{} = 1
	//iType := reflect.TypeOf(i)
	//fmt.Println(iType)
	//var d = &DoG{}
	//method := reflect.TypeOf(d).Method(0) // 获取运行时方法
	//fmt.Println(method)

	// 第二法则 从 反射到 interface{}
	//var i string = "wcg"
	//fmt.Printf("%p\n", &i)
	//i = reflect.ValueOf(i).Interface().(string)
	//fmt.Printf("%p\n", &i)

	//var i interface{} = 1
	//i = reflect.ValueOf(i).Interface().(string) // 只能转换为特定的类型,并不能随意转换
	//fmt.Println(i)

	// 第三法则	要修改反射变量,其值必须可设置
	//var i = 1
	//reflect.ValueOf(&i).Elem().Set(reflect.ValueOf(100))
	//fmt.Println(i)
}

type DucK interface {
	Like(act string)
}

type DoG struct {
	Name string
	Act  string
}

func (d *DoG) Like(act string) {
	fmt.Println(act)
}
