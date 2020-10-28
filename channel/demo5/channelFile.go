package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func writerLog(ch <-chan int, exitCh chan<- bool) {

	// 打开文件
	file, err := os.OpenFile("chanLog.log", os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)
	if err != nil {
		return
	}

	for {
		time.Sleep(time.Millisecond)
		num, ok := <-ch
		if !ok {
			break
		}
		file.WriteString(strconv.Itoa(num))
		file.WriteString("\n")

	}
	exitCh<-true
}

func writerChan(ch chan<- int) {
	for i := 1; i <= 8000; i++ {
		ch<-i
	}
	close(ch)
}

func main(){
	var ch = make(chan int, 1000)
	var exitCh = make(chan bool, 8)
	go writerChan(ch) // 写入通道

	go writerLog(ch, exitCh) // 写入文件
	time.Sleep(time.Second * 5)
	go func() {
		for i := 0; i < 8; i++ {
			<-exitCh
		}
		close(exitCh)
	}()

	fmt.Println("done")
}
