package main

import (
	"fmt"
	"io"
	"net"
)

func getmsg(conn net.Conn) {

	defer conn.Close()
	for {
		s := make([]byte, 1024)
		n, err := conn.Read(s) // 等待客户端发送信息，如果没有发送，会阻塞
		if err == io.EOF {
			fmt.Println("客户端退出")
			return
		}
		fmt.Print(string(s[:n]))
	}
}

func main() {
	// 开始监听8888
	fmt.Println("开始监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")

	if err != nil {
		fmt.Println("listen err = ", err)
	}

	defer listen.Close() // 延时关闭

	for {
		fmt.Println("等待连接....")
		conn, err := listen.Accept() // 等待连接
		if err != nil {
			fmt.Println("conn err = ", err)
		} else {
			fmt.Printf("ip ： %v 连接成功\n", conn.RemoteAddr().String()) // 获取客户端的ip
			//fmt.Printf("conn = %v \n", conn.RemoteAddr().Network()) // 获取客户端的连接方式
		}
		go getmsg(conn)
	}
}
