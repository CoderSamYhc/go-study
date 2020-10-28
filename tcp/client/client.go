package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("dial err = ", err)
		return
	}
	fmt.Printf("dial = %v \n", conn)
	reader := bufio.NewReader(os.Stdin)// 标准输入

	for {
		content, err := reader.ReadString('\n') // 获取输入的信息

		if err != nil {
			fmt.Println("read string err = ", err)
		}
		if content == "exit\n" {
			break
		}
		_, err = conn.Write([]byte(content)) // 客户端传输数据 返回 传输的字节

		if err != nil {
			fmt.Println("write err = ", err)
		}
	}

}
