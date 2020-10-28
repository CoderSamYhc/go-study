package main

import (
	"fmt"
)

func write(ch chan int) {
	for i := 0; i < 50; i++ {
		ch<-i
		fmt.Printf("write....%v \n", i)
	}
	close(ch) // 关闭
}

func read(ch chan int, flag chan bool) {

	for v := range ch {
		fmt.Printf("read....%v \n", v)
	}
	flag<-true
	close(flag)
}
func main() {
	//var c = make(chan int, 2)
	//
	//fmt.Printf("%T  %v \n", c, c)
	//
	//c<-10
	//n := 100
	//c<-n
	////c<-n // 超过长度时再进行插入 deadlock err
	//fmt.Println(len(c))
	//num1 := <-c
	//num2 := <-c
	////num3 := <-c // 全部取出后再进行取出 deadlock err
	//
	//fmt.Println(num1, num2, len(c))
	//var ch = make(chan int, 50)
	//var flag = make(chan bool, 1)
	//// 写入
	//go write(ch)
	//// 读取
	//go read(ch, flag)
	//
	//for  {
	//	_, f := <-flag
	//	if  f {
	//		break
	//	}
	//}


}
