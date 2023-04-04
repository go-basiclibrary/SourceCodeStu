package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s []string
	//marshal, err := json.Marshal(s)
	err := json.Unmarshal([]byte("null"), &s)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
