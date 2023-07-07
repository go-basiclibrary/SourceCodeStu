package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("E:\\Project\\Study\\GoStudy\\error\\jike\\countLines")
	if err != nil {
		panic(err)
	}
	lines, err := countLines(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(lines)
}

// 统计文件行数
func countLines(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	var line int
	for sc.Scan() {
		line++
	}
	return line, sc.Err()
}
