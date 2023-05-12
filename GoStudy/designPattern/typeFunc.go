package main

import (
	"fmt"
	"net/http"
)

type ConfigGenerator interface {
	CreateAct()
}

type defaultConfigGenerator struct {
	additionalCall additionalFunc
}

func (d *defaultConfigGenerator) CreateAct() {
	fmt.Println("在这里我在创建活动呢!")

	// 如果有特定的回调,那么我将进行执行回调哦
	if err := d.additionalCall; err != nil {
		fmt.Println("回调执行了哦")
		return
	}
}

// 定义一个Type Func写入一些钩子内容
type additionalFunc func(msg string) error

func NewDefaultConfigGenerator(additionalCall additionalFunc) ConfigGenerator {
	return &defaultConfigGenerator{
		additionalCall: additionalCall,
	}
}

var additionalCallMap = make(map[int]additionalFunc)

// GetAdditionalCall 这里后续如果要我改造,这里应该给一个模板类型
func GetAdditionalCall(actType int) additionalFunc {
	if call, ok := additionalCallMap[actType]; ok {
		return call
	}
	return func(msg string) error {
		fmt.Println("默认没有类型不做执行")
		return nil
	}
}

func init() {
	RegisterAdditionalCall(1, twitchDropsAdditional)
}

func twitchDropsAdditional(msg string) error {
	fmt.Println("我是自定义动作")
	return nil
}

func RegisterAdditionalCall(actType int, call additionalFunc) {
	additionalCallMap[actType] = call
}

func main() {
	//generator := NewDefaultConfigGenerator(GetAdditionalCall(1))
	//generator.CreateAct()

	resp, err := http.Get("https://baidu.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
