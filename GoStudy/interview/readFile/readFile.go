package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	ReadFile("./GoStudy/interview/readFile/eight.txt")
}

func ReadFile(path string) (int, int) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Errorf("os Open file err: %v", err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	return 0, 0
}
