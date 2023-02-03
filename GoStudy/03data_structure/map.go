package main

// Go map 运行时核心结构
//type hmap struct {
//	// 当前哈希表中的元素数量
//	count int
//	flags uint8
//	// 当前哈希表持有的buckets数量,但是因为哈希表中桶的数量都是2的倍数,所以该字段会存储对数,即len(buckets) == 2 ^ B
//	B         uint8
//	noverflow uint16
//	// 哈希表的种子,它能为哈希函数的结果引入随机性,这个值在创建哈希表时确定,并在调用哈希函数时作为参数传入
//	hash0 uint32
//
//	buckets unsafe.Pointer
//	// 是哈希表在扩容时用于保存之前buckets的字段,它的大小是当前buckets的一半
//	oldbuckets unsafe.Pointer
//	nevacuate  uintptr
//
//	extra *mapextra
//}

func main() {
	// 字面量创建map  字面量底层依然会使用make进行创建map,并通过最原始的方式向map追加元素
	//m := map[string]int{
	//	"wcg": 1,
	//}
	//fmt.Println(m)

	// var 创建map
	// 那么我的想法是,如果我们这个map是在特定条件下才会进行使用,那么我们先进行var
	// 如果在该方法内,一定会使用该数据结构,则直接使用make进行创建map,开辟内存空间
	//var m map[string]int
	//m = make(map[string]int)
	//m["1"] = 1
	//fmt.Printf("%p\n", m)

	// 运行时获取不存在的key不会造成panic,会返回该类型的零值
	//m := map[string]int{
	//	"wcg": 1,
	//}
	//i := m["hehe"]
	//fmt.Println(i)

	// value 为指针,是nil
	//m := map[string]*int{}
	//v := m["w"]
	//fmt.Println(v)
}
