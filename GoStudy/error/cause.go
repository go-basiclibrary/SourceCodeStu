package main

import (
	"fmt"
	"github.com/spf13/cast"
	"strings"
)

type taskErr struct {
	successServiceItemId   []string
	failedServiceItemSonId []string
	causeErr               []error
}

func (t *taskErr) Error() string {
	if t == nil {
		return ""
	}
	if t.causeErr == nil {
		return ""
	}
	var errString strings.Builder
	for i, err := range t.causeErr {
		errString.WriteString(fmt.Sprintf("%d %s\n", i+1, err.Error()))
	}
	return fmt.Sprintf("task err:\n%s", errString.String())
}

func main() {
	//var errs *taskErr
	//errs = &taskErr{
	//	successServiceItemId:   []string{},
	//	failedServiceItemSonId: []string{},
	//	causeErr: []error{
	//		fmt.Errorf("%s", "test err"),
	//		fmt.Errorf("%s", "test err11"),
	//	},
	//}
	//fmt.Println(errs)
	fmt.Println(cast.ToFloat32("21.0001") < 0)
	fmt.Println("1.25" < "2.01")
}
