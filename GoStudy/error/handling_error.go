package main

import (
	"GoStudy/error/util"
	"fmt"
	"git.tencent.com/trpc-go/trpc-go/errs"
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
	//err := errs.Wrapf(fmt.Errorf("%v", "cc"), 1, "abc")

	// 应用代码
	//err := errors.New("tesT stt")
	//errors.Errorf("fmt errorf : %v", errors.New("test"))

	// 和其他库协作运行
	// wrap error  只在应用内使用
	//err = errors.Wrapf(err, "failed to xx %q", "abc")
	//err = errors.Wrapf(err, "failed to xx %q", "abc")

	// 在请求入口,使用%+v把堆栈详情记录
	//err = errors.Cause(err) // 获取最初被包装的error
	//err = errors.Unwrap(err)
	//log.Errorf("%+v", err)
	//err = errs.Wrapf(err, 101, "hehe")
	////err = err.(*errs.Error).Cause()
	////fmt.Println(err)
	//err = err.(*errs.Error).Unwrap()
	//fmt.Println(err)
	//err := errors.New("abc")

	//err := fmt.Errorf("%s", "abc")
	//err = NewError(WrapError(NewErrs(1, "2"), err), "/handing_error/New")
	//err = NewErrs(100, "3")
	//err = errors.New("abc")

	//err := errs.New(1, "1")

	//err := util.WrapError(errs.New(1, "1"), fmt.Errorf("%s", "abc"))
	//fmt.Println(err)

	err := NewError(
		util.WrapError(errs.New(1, "22"),
			fmt.Errorf("%s", "abc")),
		"/1")
	fmt.Println(Code(err), Msg(err))
}

func NewErrs(code int, msg string) *errs.Error {
	return errs.New(code, msg).(*errs.Error)
}

type Error struct {
	// 接入tencent errs capacity 外部错误  用于返回报错
	WrapErr *errs.Error
	// 错误路径
	ErrPath string
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("err:%s, errpath:%s", e.WrapErr, e.ErrPath)
}

func Code(e error) int {
	if e == nil {
		return 0
	}
	err, ok := e.(*Error)
	if !ok {
		if _, ok = e.(*errs.Error); ok {
			return errs.Code(e)
		}
		return -1
	}
	return int(err.WrapErr.Code)
}

func Msg(e error) string {
	if e == nil {
		return errs.Success
	}
	err, ok := e.(*Error)
	if !ok {
		if tErr, ok := e.(*errs.Error); ok {
			return tErr.Msg + tErr.Cause().Error()
		}
		return e.Error()
	}
	return err.WrapErr.Msg
}

func NewError(err *errs.Error, path string) error {
	return &Error{
		WrapErr: err,
		ErrPath: path,
	}
}

//func (c *Error) Error() string {
//	if c == nil {
//		return ""
//	}
//	return fmt.Sprintf("external err:%s, errpath:%s",
//		c.External.Error(), c.ErrPath)
//}

func WrapError(errExternal *errs.Error, errInternal error) *errs.Error {
	//errMsg := fmt.Sprintf("%s:%v", errs.Msg(errExternal), errInternal)
	//return errs.New(errs.Code(errExternal), errMsg)
	return errs.Wrap(errInternal, errs.Code(errExternal), errs.Msg(errExternal)).(*errs.Error)
}
