package main

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

var sqlNoRows = errors.New("no rows")

func main() {
	i, err := service(nil)
	if err != nil {
		is := errors.Is(err, sqlNoRows)
		if is {
			fmt.Println(err)
			fmt.Println(http.StatusNotFound)
			return
		}
		panic(err)
	}
	fmt.Println(i)
}

func getUser(id int64) (interface{}, error) {
	return nil, sqlNoRows
}

func service(req interface{}) (interface{}, error) {
	// service logic
	if req != nil {
		return nil, nil
	}

	user, err := getUser(1)
	if err != nil {
		//return nil, errors.Wrap(err, "get user err")
		return nil, fmt.Errorf("%v %w", "get user err", err)
	}

	return user, err
}
