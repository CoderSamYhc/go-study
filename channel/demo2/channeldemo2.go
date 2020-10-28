package main

import (
	"fmt"
	"time"
)

func writer(ch chan int) {
	for i := 1; i <= 8000; i++ {
		ch<-i
		//fmt.Printf("writer....%v\n", i)
	}
	close(ch)
}

func reader(ch chan int, pre chan int, flag chan bool) {

	f := false
	for  {
		time.Sleep(time.Millisecond)
		v, err := <-ch
		if !err {
			break
		}
		//fmt.Printf("read....%v \n", v)
		f = true

		for i := 2; i < v; i++ {
			if v%i == 0 {
				f = false
				break
			}
		}
		if f {
			pre<-v
		}
	}


	flag<-true
}

func main() {
	var ch = make(chan int, 1000)
	var pre = make(chan int, 2000)
	var flag = make(chan bool, 4)

	go writer(ch)

	for i := 1; i <= 4; i++ {
		go reader(ch, pre, flag)
	}

	go func(){
		for i := 0; i < 4; i++ {
			<-flag
		}

		close(pre)
	}()

	for  {
		res, err := <-pre
		if !err {
			break
		}

		fmt.Println(res)
	}
	fmt.Println("stop")
}
