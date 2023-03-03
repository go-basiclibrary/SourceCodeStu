package main

import "fmt"

// for & range
func main() {
	//var s = []int{1, 2, 3}
	//// 是否会永动循环
	//for _, v := range s {
	//	s = append(s, v)
	//}
	//// [1 2 3 1 2 3]  循环具体能走几步,和初始循环的条件有关
	//fmt.Println(s)

	// for 中常见赋值问题
	//var s = []int{1, 2, 3}
	//var newArr []*int
	//for _, v := range s {
	//	// 这里会有问题,v最终会变成3
	//	newArr = append(newArr, &v)
	//}
	//fmt.Println(newArr)

	// for range的三种遍历方式
	//var newArr []int
	//for range newArr { // 没什么意义
	//}
	//for i, _ := range newArr {
	//	// 只要索引
	//}
	//for i, v := range newArr {
	// 既要索引,又要v
	// 但是v每次会被newArr[i]变量覆盖,并且是发生了复制
	// 所以大切片,在for时是可以优化性能的
	//}

	// for range 修改其中的值  无法修改哦
	//newSlice := []int{1, 2, 3}
	//for _, v := range newSlice {
	//	v = 5
	//	v = 6
	//	v = v
	//}
	//fmt.Println(newSlice)

	// for range 修改指针  指针是可以修改值的哦
	//var newS []*int
	//var i int
	//newS = append(newS, &i)
	//for _, v := range newS {
	//	i = 100
	//	v = &i
	//	fmt.Println(v)
	//}
	//for _, v := range newS {
	//	fmt.Println(*v)
	//}

	//slice := new([]int)
	//fmt.Println(slice == nil)
	//fmt.Println(*slice == nil)

	//var ch = make(chan int, 10)
	//
	//go func() {
	//	for {
	//		time.Sleep(1e9)
	//		ch <- 12
	//		close(ch)
	//	}
	//}()
	//for v := range ch {
	//	fmt.Println(v)
	//}

	var i = 1
	i += 1
	fmt.Println(i)
}
