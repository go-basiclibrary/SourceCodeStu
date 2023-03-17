package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

func main() {
	ginPanic()
	//跨协程失效
	//CrossCtrip()

	// 嵌套崩溃
	//NestedCrash()

	// Go exit功能
	//GoExit()

	// 恢复后的功能
	//GoPanicRecover()

	// 查看defer recover会处理到哪一步
	//DeferPanic()
	//fmt.Println("abc")

	// go 嵌套panic
	//PanicQt()
}

func PanicQt() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		panic("456")
	}()
	panic("123")
}

func DeferPanic() error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("默默阿达")
	return nil
}

func GoPanicRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer func() {
		fmt.Println("abc")
	}()
	var i int

	i++
	panic(i)
	fmt.Printf("this is %d\n", i)
}

func GoExit() {
	go func() {
		go func() {
			time.Sleep(1e9)
			fmt.Println("hehe")
		}()
		panic("abc")
	}()
	go func() {
		time.Sleep(1e9)
		fmt.Println("123")
	}()
	fmt.Println("enen")
	time.Sleep(2e9)
}

func NestedCrash() {
	defer fmt.Println("abc")
	defer func() {
		defer func() {
			panic("1")
		}()
		panic("2")
	}()
	panic("3")
}

func CrossCtrip() {
	// 外部defer语句没有执行
	defer println("in main")
	go func() {
		defer println("in goroutine")
		panic("")
	}()
	go func() {
		// 这个是可能不被执行的
		defer println("in goroutine2")
		time.Sleep(0.1e9)
	}()

	time.Sleep(1e9)
}

func ginPanic() {
	engine := gin.Default()
	g := engine.Group("/test")
	g.GET("/tt", func(context *gin.Context) {
		//go func() {
		//	defer func() {
		//		if err := recover(); err != nil {
		//			fmt.Println(err)
		//		}
		//	}()
		//	panic(123)
		//}()
		var once sync.Once
		for i := 0; i < 3; i++ {
			once.Do(func() {
				fmt.Fprintf(context.Writer, "abc+%d", i)
			})
		}
		once.Do(func() {
			fmt.Fprintf(context.Writer, "abc")
		})
	})

	engine.Run(":8080")
}
