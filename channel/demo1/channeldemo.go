package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Person struct {
	Name string
	Age int
	Address string
}

func main() {
	//var ch = make(chan interface{}, 10)
	//ch<-Person{"Sam", 19}
	//ch<-Person{"Kobe", 19}
	//ch<-10
	//
	//sam := <-ch
	//fmt.Printf("%T   %v \n", sam, sam)
	//fmt.Printf("%T   %v \n", sam, sam.(Person).Name)

	var ch = make(chan Person, 10)

	for i := 0; i < 10; i++ {
		p := Person{
			Name : "Sam" + strconv.Itoa(i),
			Age: rand.Intn(100),
			Address: "Address" + strconv.Itoa(i),
		}
		ch<-p
	}
	close(ch) // 遍历前需要关闭管道，否则报 deadlock
	for v := range ch {
		fmt.Println(v)
	}
}
