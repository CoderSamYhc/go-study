package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)

	// lock 是全局互斥锁
	// sync 是包 同步
	// Mutex 互斥

	lock sync.Mutex
)

func tests(n int) {

	res := 1
	for i := 1; i <=n; i++ {
		res *= i
	}
	lock.Lock() // 加锁
	myMap[n] = res // concurrent map writes
	lock.Unlock() // 解锁
}

func main() {
	for x := 1; x <= 200; x++ {
		go tests(x)
	}
	// 防止主线程结束时，协程未结束
	time.Sleep(time.Second * 20)

	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]%d \n", i, v)
	}
	lock.Unlock()
}

