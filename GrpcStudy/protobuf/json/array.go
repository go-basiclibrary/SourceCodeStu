package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	tags := []string{"wS", "Sw", "test"}
	marshal, err := json.Marshal(tags)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))
}
