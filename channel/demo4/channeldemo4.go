package main

import (
	"fmt"
	"time"
)

// 1.将1-2000 放入通道
// 2。起8个协程从通道取，并计算 1+。。+n的值，存放在另一个通道
// 3.8个协程完成工作后，遍历：如res[1]=1 res[2]=3....res[10]=55....res[2000]=2001000

func initChan (ch chan<- int) {
	for i := 1; i <= 2000; i++ {
		ch<-i
	}
	close(ch)
}

func sumChan(initCh <-chan int, res chan<- int, flag chan bool, id int) {
	var n int
	for {
		time.Sleep(time.Millisecond)
		num, err := <-initCh
		if !err {
			break
		}
		n = num
		var resNum int
		for i := 1; i <= num; i++{
			resNum += i
		}
		res<-resNum
	}
	fmt.Printf("worker:%d---last:%d \n", id, n)
	flag <- true
}

func main (){
	init := make(chan int, 1000)
	res := make(chan int, 2000)
	flag := make(chan bool, 8)

	go initChan(init) // 协程写入

	for i := 0; i < 8; i++ {
		go sumChan(init, res, flag, i)
	}

	go func() {
		for i := 0; i < 8; i++ {
			<-flag
		}
		close(res)
	}()
	i := 1
	for v := range res {
		fmt.Printf("res[%d]=%d \n", i, v)
		i++
	}
}
