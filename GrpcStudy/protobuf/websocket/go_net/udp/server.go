package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8000")
	if err != nil {
		log.Panic(err)
	}

	listen, err := net.ListenUDP("udp", addr)
	defer listen.Close()
	if err != nil {
		log.Panic(err)
	}

	// handle
	go handleConnection(listen)
	signal := make(chan os.Signal)
	<-signal
	// close conn
	return
}

func handleConnection(conn *net.UDPConn) {
	//recv msg
	for {
		// 设置消息长度为1024比特
		buf := make([]byte, 1024)
		// 读取消息，UDP不是面向连接的因此不需要等待连接
		length, udpAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("Error: ", err)
			continue
		}
		fmt.Println("[ server ]# UdpAddr: ", udpAddr, "Data: ", string(buf[:length]))
	}
}
