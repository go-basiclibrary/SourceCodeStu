package str

type Stu struct {
	name string
	age  int
}

func CreateStu(name string, age int) Stu {
	return Stu{name: name, age: age}
}
