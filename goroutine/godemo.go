package goroutine

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer fmt.Println("defer", i)
			if i == 5 {
				//runtime.Gosched() // 表示 5不会第一个输出
				runtime.Goexit() // 表示 终止当前协程
			}
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second)
}
