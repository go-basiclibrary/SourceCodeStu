package main

import (
	"fmt"
	"io"
	"os"
)

type header struct {
	key, value string
}

type status struct {
	code   int
	reason string
}

// 写入响应  过多的err!=nil 如何优化
func writeResponse(w io.Writer, st status, headers []header, body io.Reader) error {
	_, err := fmt.Fprintf(w, "HTTP/1.1 %d %s\r\n", st.code, st.reason)
	if err != nil {
		return err
	}
	for _, h := range headers {
		_, err = fmt.Fprintf(w, "%s: %s\r\n", h.key, h.value)
		if err != nil {
			return err
		}
	}

	if _, err = fmt.Fprintf(w, "\r\n"); err != nil {
		return err
	}
	_, err = io.Copy(w, body)
	return err
}

func main() {
	stdout := os.Stdout
	stdin := os.Stdin
	err := writeResponse1(stdout, status{
		code:   404,
		reason: "这个文件找不到",
	}, []header{
		{
			key:   "x-user",
			value: "ws",
		},
		{
			key:   "x-seq",
			value: "170699829123",
		},
	}, stdin)
	if err != nil {
		panic(err)
	}
}

type errWriter struct {
	io.Writer
	err error
}

func (e *errWriter) Write(buf []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}
	var n int
	n, e.err = e.Writer.Write(buf)
	return n, e.err
}

// 优化writeResponse  总结来说就是分层做了处理,对err做了提出
func writeResponse1(w io.Writer, st status, headers []header, body io.Reader) error {
	ew := &errWriter{Writer: w}
	fmt.Fprintf(ew, "HTTP/1.1 %d %s\r\n", st.code, st.reason)
	for _, h := range headers {
		fmt.Fprintf(ew, "%s: %s\r\n", h.key, h.value)
	}
	fmt.Fprint(ew, "\r\n")
	io.Copy(ew, body)

	return ew.err
}
