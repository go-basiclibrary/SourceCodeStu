package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	g, _ := errgroup.WithContext(context.TODO())
	g.Go(func() error {
		fmt.Println("2")
		return fmt.Errorf("123")
	})
	g.Go(func() error {
		fmt.Println("2")
		return fmt.Errorf("321")
	})
	g.Go(func() error {
		fmt.Println("2")
		return fmt.Errorf("888")
	})

	g.Go(func() error {
		fmt.Println("1")
		return nil
	})

	err := g.Wait()
	fmt.Println(err)
}
