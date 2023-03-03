package main

import (
	"errors"
	"fmt"
)

// Error types 方式 os.PathError

type MyError struct {
	Msg  string
	File string
	Line int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s:%d: %s", e.File, e.Line, e.Msg)
}

func test() error {
	return &MyError{"Something happened", "server.go", 42}
}

// Opaque errors  少数情况下,这种错误处理方式肯定是不够的,与进程外的世界进行交互,需要调用方调查错误的性质,以
// 确定重试该操作是否合理
func fn() error {
	var err error
	err = errors.New("fn")
	if err != nil {
		return err
	}
	return nil
}

// 断言行为错误,而不是类型错误
type temporary interface {
	Temporary() bool
}

type temError struct {
}

func (tE *temError) Temporary() bool {
	return true
}

func (tE *temError) Error() string {
	return "abc"
}

// IsTemporary 内部提供Is方法供外部判断特殊error
func IsTemporary(err error) bool {
	te, ok := err.(temporary)
	return ok && te.Temporary()
}

func main() {
	//err := test()
	//switch err := err.(type) { // 依然将外部与内部类型进行了强耦合
	//case nil:
	//case *MyError:
	//	fmt.Println("error occurred on line:", err.Line)
	//default:
	//
	//}

	tE := &temError{}
	if IsTemporary(tE) {
		fmt.Println("这种情况下,调用方知道可以做一些其他事情")
	}
}
