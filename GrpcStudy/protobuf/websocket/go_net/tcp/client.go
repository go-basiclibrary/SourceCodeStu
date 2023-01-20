package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	tcpConn, err := net.DialTCP("tcp", nil, addr)
	defer tcpConn.Close()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		// write
		fmt.Print("[ random_w ]# ")
		bytes, _, _ := reader.ReadLine()
		tcpConn.Write(bytes)

		rb := make([]byte, 1024)
		// read
		read, _ := tcpConn.Read(rb)
		fmt.Println(string(rb[:read]))

		break
	}
}
