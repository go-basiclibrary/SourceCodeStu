package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var sig = make(chan os.Signal)

func main() {
	RemoteAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8001")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", LocalAddr, RemoteAddr)
	if err != nil {
		panic(err)
	}

	go HandleConnectionForClient(conn)
	<-sig
	return
}

// HandleConnectionForClient 读取数据, 在这里我们可以编写自己的交互程序
func HandleConnectionForClient(conn net.Conn) {
	// 监控系统信号
	go signalMonitor(conn)
	// 初始化一个缓存区
	Stdin := bufio.NewReader(os.Stdin)
	for {
		// 读取用户输入的信息，遇到换行符结束。
		fmt.Print("[ random_w ]# ")
		input, err := Stdin.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		// 删除字符串前后的空格，主要是删除换行符。
		input = strings.TrimSpace(input)
		// 空行不做处理
		if len(input) == 0 {
			continue
		}
		// 是否接收到退出指令
		switch input {
		case "quit", "exit":
			sig <- syscall.SIGQUIT
		default:
			// 发送消息给服务端
			sendMsgToServer(conn, input)
		}
	}
}

// sendMsgToServer 发送消息给服务端
func sendMsgToServer(conn net.Conn, msg string) {
	for {
		_, err := conn.Write([]byte(msg))
		if err == nil {
			break
		}
	}
}

// signalMonitor 监听系统信号，如果程序收到退出到的信号通过 Goroutine 通知 server 端，关闭连接后退出。
func signalMonitor(conn net.Conn) {

	signal.Notify(sig, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGINT)
	// 接收到结束信号退出此程序
	select {
	case <-sig:
		// 通知服务端断开连接
		_, _ = conn.Write([]byte("exit"))
		fmt.Println("\nGood Bye !!!!!")
		os.Exit(0)
	}
}
