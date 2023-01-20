package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// go net包 tcp实现长链接
func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	lis, err := net.ListenTCP("tcp", addr)
	defer lis.Close()
	if err != nil {
		panic(err)
	}
	for {
		tcpConn, err := lis.AcceptTCP() //获取tcp连接
		if err != nil {
			log.Printf("accept tcp error %v\n", err.Error())
			return
		}
		go handleConnection(tcpConn)
	}
}

func handleConnection(conn *net.TCPConn) {
	defer conn.Close()
	for {
		// recv
		b := make([]byte, 1024)
		read, err := conn.Read(b)
		if err == io.ErrClosedPipe {
			fmt.Println("EOF")
			return
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("conn read error:%v\n", err.Error())
			break
		}
		fmt.Println("[ server ]# UdpAddr: ", conn.RemoteAddr().String(), "Data: "+string(b[0:read]))

		// send  server端推送消息
		conn.Write([]byte("i'm coming"))
	}
}
