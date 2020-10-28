package main

import (
	"fmt"
	"time"
)

func main() {
	//var ch = make(chan int, 10)
	//for i := 0; i < 10; i++ {
	//	ch<-i
	//}
	//
	//for  {
	//	select {
	//		case v := <-ch:
	//			fmt.Println(v)
	//			time.Sleep(time.Millisecond)
	//		default:
	//			fmt.Println("stop")
	//			return
	//	}
	//}

	ticker := time.NewTicker(time.Millisecond * 500)

	stoper := time.NewTimer(time.Second * 2)
	for {
		select {
		case <-ticker.C:
			fmt.Println("计时器到了")
		case <-stoper.C:
			fmt.Println("结束！")
			return
		}
	}
	fmt.Println("done")

}
