package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"
)

func test(id int) {
	fmt.Println("开始写入")
	f, _ := os.OpenFile("ids.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	defer f.Close()
	var stringBuilder bytes.Buffer
	stringBuilder.WriteString("---start-- \n")
	stringBuilder.WriteString(strconv.Itoa(id))
	stringBuilder.WriteString("\n")
	stringBuilder.WriteString("---end-- \n")
	str := stringBuilder.String()
	fmt.Println(str)
	res, err := f.WriteString(str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("结束写入", res)

}

func main() {

	for i := 1; i <= 10; i++ {
		go test(i)
	}
	time.Sleep(time.Second) // 防止主线程结束时，协程未结束
}