package main

import "fmt"

type User struct {
	Name string
	Age  int
	Size int
}

func main() {
	u := &User{
		Name: "WangShao",
		Age:  10,
		Size: 1,
	}
	ChangeUser(u)
	fmt.Println(u)
}

func ChangeUser(u *User) {
	u.Size = 18
}
