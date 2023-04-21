package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Slice []int

func (A Slice) Append(value int) {
	A1 := append(A, value)

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&A))
	fmt.Printf("A Data:%d,Len:%d,Cap:%d\n", sh.Data, sh.Len, sh.Cap)

	sh1 := (*reflect.SliceHeader)(unsafe.Pointer(&A1))
	fmt.Printf("A1 Data:%d,Len:%d,Cap:%d\n", sh1.Data, sh1.Len, sh1.Cap)
}

// 数据结构之--切片
func main() {
	// 切片底层数组扩容是地址是否会发生变化
	//ints := make([]int, 0, 3)
	//ints = append(ints, 1)
	//fmt.Printf("%p\n", ints)
	//ints = append(ints, 2, 3)
	//fmt.Printf("%p\n", ints)

	// 空切片是否为nil
	//ints := []int{}
	//fmt.Println(ints == nil)

	// var 创建引用类型是否为nil
	//var ints []int
	//fmt.Println(ints == nil)
	//fmt.Printf("%p\n", ints)

	// 修改原数组是否会改变数据
	// 是否会修改原切片或者数据的数据,就看他们底层的数组指向是否发生改变
	//var arr = [3]int{1, 2, 3}
	//ints := arr[:]
	//arr[0] = 101
	//fmt.Println(ints)
	//ints[0] = 1
	//fmt.Println(arr)
	//ints = append(ints, 10)
	//ints[0] = 101
	//fmt.Println(arr, ints)

	// copy切片,修改是否会影响原切片  woC!!不会!
	//ints := []int{1, 2, 3}
	//var newInts = make([]int, 3)
	//copy(newInts, ints)
	//newInts[0] = 100
	//fmt.Println(ints)

	// Append
	//mSlice := make(Slice, 10, 20)
	//mSlice.Append(5)
	//fmt.Println(mSlice)

	// 内存对齐
	//arr := make([]int, 0)
	//arr = append(arr, 1, 2, 3, 4, 5)
	//fmt.Println(len(arr), cap(arr))

	// 字面量创建切片
	//slice := []int{}[:]
	//slice = []int{}
	//slice = []int{1, 2, 3}[:]
	//var arr [3]int
	//slice := arr[:]
	//fmt.Println(slice)
	//var slice = []int{1, 2, 3, 4, 5}
	//fmt.Println(len(slice), cap(slice))

	//var b = make([]byte, 0, 1)
	//generateB(b)
	//fmt.Println(b)

	// 如果底层没有进行扩容,那么内部修改值则会影响外部
	// 如果底层切片引用变了,那么内部修改不会影响外部
	//res := []int{1}
	//changeSlice(res)
	//fmt.Println(res)

	// slice append ptr is change
	//res := []int{1, 2, 3}
	//fmt.Printf("%p\n", res)
	//
	//res = res[:2]
	//fmt.Printf("%p\n", res)
	//
	//res = append(res, 1)
	//fmt.Printf("%p\n", res)
	//fmt.Println(res)
	//
	//res = append(res, 1)
	//fmt.Printf("%p\n", res)
	//fmt.Println(res)

	//var res []int
	//fmt.Println(res == nil)

	//var res = make([]int, 0)
	//for i := 0; i < 10; i++ {
	//	res = append(res, i)
	//	fmt.Println(cap(res))
	//}
	//res = append(res, 1, 2, 3, 4)
	//fmt.Println(cap(res))
	//fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&res)))

	// deep copy
	//var res = make([]int, 2)
	//copy(res, []int{1, 2, 3})
	//fmt.Println(res)
}

func changeSlice(res []int) {
	//res = append(res, 1000)
}

func generateB(b []byte) {
	b = append(b, 15)
}
