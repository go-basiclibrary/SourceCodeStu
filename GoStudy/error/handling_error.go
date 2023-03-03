package main

import (
	"fmt"
	"github.com/pkg/errors"
)

// Handling error
func main() {
	//errs.SetTraceable(true)
	//_, err := util.IntToString(5)
	//if errors.Is(err, util.ErrNotFound) {
	//	fmt.Printf("the err is not found! %+v\n", err)
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(err)
	//}

	// 应用代码
	err := errors.New("tesT stt")
	errors.Errorf("fmt errorf : %v", errors.New("test"))

	// 和其他库协作运行
	// wrap error  只在应用内使用
	err = errors.Wrapf(err, "failed to xx %q", "abc")
	err = errors.Wrapf(err, "failed to xx %q", "abc")

	// 在请求入口,使用%+v把堆栈详情记录
	err = errors.Cause(err) // 获取最初被包装的error
	fmt.Println(err)
}
