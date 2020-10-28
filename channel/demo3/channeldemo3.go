package main

import (
	"fmt"
	"time"
)

func createWorker(id int) chan<- int { // 外部只能增加通道数据
	ch := make(chan int)
	go func() { // 开启协程
		for {
			fmt.Printf("workerId:%d, %c \n", id, <-ch)
		}
	}()
	return ch
}

func main() {
	var channels [10]chan<- int // 需要对应
	for i := 0; i < 10; i++{
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}
