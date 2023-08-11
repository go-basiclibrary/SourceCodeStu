package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	open, err := os.Open("E:\\Project\\Study\\GoStudy\\jike\\02\\errhandling\\new")
	if err != nil {
		panic(err)
	}
	//line, err := countLines(open)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(line)

	lines02, err := countLines02(open)
	if err != nil {
		panic(err)
	}
	fmt.Println(lines02)
}

func countLines(r io.Reader) (int, error) {
	var (
		br    = bufio.NewReader(r)
		lines int
		err   error
	)

	for {
		s, err := br.ReadString('\n')
		fmt.Println(s)
		lines++
		if err != nil {
			break
		}
	}

	if err != io.EOF && err != nil {
		return 0, err
	}
	return lines, nil
}

// 改进版countLines
func countLines02(r io.Reader) (int, error) {
	var (
		sc    = bufio.NewScanner(r)
		lines int
	)

	for sc.Scan() {
		lines++
	}

	return lines, sc.Err()
}
