package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.TODO(), "1", "val1")
	ctx = context.WithValue(ctx, "1", "val2")

	value := ctx.Value("1")
	fmt.Println(value)
}
