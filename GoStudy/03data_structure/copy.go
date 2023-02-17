package main

type User struct {
	Name string
	Age  int
	Size int
}

func (u *User) GetName() {
	println(u)
}

func main() {
	u := &User{
		Name: "WangShao",
		Age:  10,
		Size: 1,
	}
	ChangeUser(u)
	println(u)
	u.GetName()
}

func ChangeUser(u *User) {
	println(u)
	u.Size = 18
}
