package main

import (
	"fmt"
)

// 反射
func main() {
	//var d = &DoG{}
	// 通过反射必须要这么才能判断是否实现了该接口
	// 这里Implements后面必须填接口类型,否则运行时panic
	//bo := reflect.TypeOf(d).Implements(reflect.TypeOf((*Bigger)(nil)).Elem())
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

	//var d DucK = &DoG{}
	//s := reflect.TypeOf(d).String()
	//fmt.Println(s)

	//使用反射更新struct
	//var stu = str.Stu{}
	//sT := str.CreateStu("wcg", 5)
	//reflect.ValueOf(&stu).Elem().Set(reflect.ValueOf(sT))
	//fmt.Println(stu)
	//var stuZ = &str.Stu{}
	//sTD := str.CreateStu("wcg", 5)
	//reflect.ValueOf(&stuZ).Elem().Set(reflect.ValueOf(&sTD))
	//fmt.Println(stuZ)

	// 使用反射调用执行方法
	//v := reflect.ValueOf(Add)
	//if v.Kind() != reflect.Func {
	//	return
	//}
	//t := v.Type()
	//argv := make([]reflect.Value, t.NumIn())
	//for i := range argv {
	//	if t.In(i).Kind() != reflect.Int {
	//		return
	//	}
	//	argv[i] = reflect.ValueOf(i) // 传入参数
	//}
	//result := v.Call(argv) // 调用函数
	//if len(result) != 1 || result[0].Kind() != reflect.Int {
	//	return
	//}
	//fmt.Println(result)
	//fmt.Println(result[0].Int())
}

func Add(a, b int) int {
	return a + b
}

type DucK interface {
	Like(act string)
}

type Bigger interface {
	Big()
}

type DoG struct {
	Name string
	Act  string
}

func (d *DoG) Like(act string) {
	fmt.Println(act)
}
