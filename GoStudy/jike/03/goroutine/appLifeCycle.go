package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	var a App
	a.track.Event("event")
	ctx, _ := context.WithTimeout(context.TODO(), 1*time.Second)

	err := a.track.ShutDown(ctx)
	if err != nil {
		log.Println(err)
	}
}

type Tracker struct {
	// 接入wg,用于追踪每一个goroutine的创建
	wg sync.WaitGroup
}

func (t *Tracker) Event(data string) {
	t.wg.Add(1)

	go func() {
		defer t.wg.Done()
		time.Sleep(time.Millisecond)
		log.Println(data)
	}()
}

// ShutDown
// 虽然该方法控制了所有该控制的周期
// 使用了wg控制g,使用ctx控制time
// 但是该内容大量的创建g来处理任务,代价是高的
func (t *Tracker) ShutDown(ctx context.Context) error {
	ch := make(chan struct{})
	go func() {
		t.wg.Wait()
		close(ch)
	}()

	select {
	case <-ch:
		return nil
	case <-ctx.Done():
		return errors.New("timeout")
	}

	//// 使用wg追踪了生命周期,但是否需要考虑有任务泄露问题
	//t.wg.Wait()
}

type App struct {
	track Tracker
}

func (a *App) Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	// 没有管理该goroutine的生命周期
	// 无法保证创建的goroutine生命周期管理,会导致最常见的问题,就是服务关闭时候,有一些时间丢失
	//go a.track.Event("this event")

	a.track.Event("this event")
}
