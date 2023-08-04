package main

import (
	"errors"
	"fmt"
)

func check(n int) {
	res, err := positive(n)
	if err != nil {
		fmt.Println(n, err)
		return
	}
	if res {
		fmt.Printf("%d is positive\n", n)
	} else {
		fmt.Printf("%d is negative\n", n)
	}
}

func main() {
	check(-1)
	check(0)
	check(1)
}

func positive(n int) (bool, error) {
	if n == 0 {
		return false, errors.New("undefined")
	}
	return n > -1, nil
}
