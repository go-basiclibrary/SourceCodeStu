package main

import (
	"crypto/md5"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	done := make(chan struct{})
	paths := make(chan string)
	c := make(chan result)
	var wg sync.WaitGroup
	const numDigesters = 20
	wg.Add(numDigesters)

	for i := 0; i < numDigesters; i++ {
		// 控制20个数量的协程来使用该操作
		go func() {
			digester(done, paths, c)
			// 任务完成
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(c)
		close(done)
		close(paths)
	}()
}

type result struct {
	path string
	md5  [16]byte
	err  error
}

func digester(done <-chan struct{}, paths <-chan string, c chan<- result) {
	for path := range paths {
		data, err := os.ReadFile(path)
		select {
		case c <- result{path, md5.Sum(data), err}:
		case <-done:
			return
		}
	}
}

// 并发读写文件内容
func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)
	go func() {
		defer close(paths)
		errc <- filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- path:
			case <-done:
				return errors.New("walk canceled")
			}
			return nil
		})
	}()

	return paths, errc
}
