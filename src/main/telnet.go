package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func server(address string, exitChan chan int) {
	l, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err.Error())
		exitChan <- 1
	}
	fmt.Println("listen:" + address)
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		handleSession(conn, exitChan)
	}
}

func handleSession(conn net.Conn, exitChan chan int) {
	fmt.Println("Session started:")
	reader := bufio.NewReader(conn)
	for {
		s, err := reader.ReadString('\n')
		if err == nil {
			s = strings.TrimSpace(s)

			if !processTelnetCommand(s, exitChan) {
				conn.Close()
				break
			}

			conn.Write([]byte(s + "\r\n"))
		} else {
			fmt.Println("Session closed")
			conn.Close()
			break
		}
	}
}

func processTelnetCommand(str string, exitChan chan int) bool {
	// @close指令表示终止本次会话
	if strings.HasPrefix(str, "@close") {
		fmt.Println("Session closed")
		// 告诉外部需要断开连接
		return false
		// @shutdown指令表示终止服务进程
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("Server shutdown")
		// 往通道中写入0, 阻塞等待接收方处理
		exitChan <- 0
		// 告诉外部需要断开连接
		return false
	}
	// 打印输入的字符串
	fmt.Println(str)
	return true
}

func telnetServer() {
	// 创建一个程序结束码的通道
	exitChan := make(chan int)
	// 将服务器并发运行
	go server("127.0.0.1:7001", exitChan)
	// 通道阻塞, 等待接收返回值
	code := <-exitChan
	// 标记程序返回值并退出
	os.Exit(code)
}
