package util

import (
	"errors"
	"fmt"
	"git.tencent.com/trpc-go/trpc-go/errs"
	"github.com/spf13/cast"
)

var ErrNotFound = errors.New("not found")
var DBErr = errs.New(1001, "db error")

func IntToString(i int) (string, error) {
	err := WrapError(DBErr, fmt.Errorf("find act by err: %w", ErrNotFound))
	//err = errors.Unwrap(err)
	return cast.ToString(i), err
}

func WrapError(errExternal, errInternal error) error {
	errMsg := fmt.Sprintf("%s", errs.Msg(errExternal))
	return errs.Wrap(errInternal, errs.Code(errExternal), errMsg)
}
