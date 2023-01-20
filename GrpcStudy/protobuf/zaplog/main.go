package main

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func init() {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./test.log",
		"stderr",
	}
	Logger, err := cfg.Build()
	if err != nil {
		panic("zap.NewDevelopment err:" + err.Error())
	}
	zap.ReplaceGlobals(Logger) // 将logger注册到全局中  后续使用直接使用zap.S() or zap.L() 即可
}

func main() {
	zap.S().Infof("test %v", "i'm test")
	zap.L().Info("这里有问题?")
	//runtime.GOMAXPROCS(1)
	//logger, err := NewLogger()
	//if err != nil {
	//	panic(err)
	//}
	//
	//sugarLogger := logger.Sugar()
	//sugarLogger.Infow("123", "456", "789")

	//int_chan := make(chan int, 1)
	//string_chan := make(chan string, 1)
	//
	//int_chan <- 1
	//string_chan <- "abc"
	//select {
	//case value := <-int_chan:
	//	fmt.Println(value)
	//case value := <-string_chan:
	//	fmt.Println(value)
	//}
}
