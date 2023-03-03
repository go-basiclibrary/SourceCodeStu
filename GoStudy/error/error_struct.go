package main

import (
	"errors"
	"fmt"
)

// Create a named type for our new error type.
type errorString string

// Implement the error interface
func (e errorString) Error() string {
	return string(e)
}

// New creates interface values of type error.
func New(text string) error {
	return errorString(text)
}

var ErrNamedType = New("EOF")
var ErrStructType = errors.New("EOF")

func main() {
	if ErrNamedType == New("EOF") { // 我们不希望,一个string和我们相比,就是我们自己
		fmt.Println("Named Type Error")
	}

	if ErrStructType == errors.New("EOF") {
		fmt.Println("Struct Type Error")
	}

	// 常规项目中多使用该方式来判断error是否是这个错误
	if ErrStructType == Err() { // 我们希望,当我们返回了一个我们自定义的错误,你只能用我产生错误的结构来判断是不是我定义的错误
		fmt.Println("Struct Type True")
	}
}

func Err() error {
	return ErrStructType
}
