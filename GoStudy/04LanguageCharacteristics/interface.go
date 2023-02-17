package main

// Duck 04 interface
type Duck interface {
	Quick()
	Build()
}

type Cat struct {
	Name string
}

func (c Cat) Quick() {
	c.Name = "测试"
}

func (c Cat) Build() {

}

func main() {
	//var d Duck = Cat{}  这里会报错
	//var dZ Duck = &Cat{}

	//var d Duck = Cat{}
	//d.Quick()
	//fmt.Println(d.(Cat).Name)

	//var d Duck = &Cat{}
	//d.Quick()
	//fmt.Println(d.(*Cat).Name)

	//var c Cat
	//c.Build()
	//fmt.Println(c.Name)

	//var c *Cat
	//fmt.Println(c == nil)
	//fmt.Println(NilOrNot(c))

	// interface类型断言,以及接口类型断言尽量使用switch,case
	//var d interface{} = Cat{}
	//d.(*Cat).Quick()
	//var d Duck = Cat{}
	//d.(*Cat).Quick()

	//var c Duck = Cat{}
	//t := c.(*Dog) // 编译时进行判定
	//var c interface{} = Cat{}
	//dog := c.(*Dog) // 运行时进行判定
	//fmt.Println(dog)

	//var c Duck = Cat{}
	//c.(*Dog).Build()
}

type Dog struct {
}

func (d *Dog) Quick() {
}

func (d *Dog) Build() {
}

func NilOrNot(v interface{}) bool {
	// 当这里我们去拿内部的类型时,依然是true
	// 相当于我们做了一个新的变量,包含了原有的变量和类型
	return v.(*Cat) == nil
}
