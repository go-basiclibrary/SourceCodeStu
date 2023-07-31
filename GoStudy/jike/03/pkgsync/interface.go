package main

import "fmt"

// data race interface
// interface底层结构是 two machine world
func main() {
	var ben1 = &ben{id: 10, name: "ben"}
	var jerry1 = &jerry{name: "jerry"}
	var marker IceCreamMaker = ben1

	var loop0, loop1 func()
	loop0 = func() {
		marker = ben1
		go loop1()
	}
	loop1 = func() {
		marker = jerry1
		go loop0()
	}
	go loop0()

	for {
		marker.Hello()
	}
}

type IceCreamMaker interface {
	Hello()
}

type ben struct {
	id   int
	name string
}

func (b *ben) Hello() {
	fmt.Printf("ben says,\"Hello my name is %s\"\n", b.name)
}

type jerry struct {
	name string
}

func (j *jerry) Hello() {
	fmt.Printf("jerry says,\"Hello my name is %s\"\n", j.name)
}
