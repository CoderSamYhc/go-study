package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func main () {

	// 创建输出文件
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("create cpu.prof err", err)
		return
	}

	// 获取系统信息
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("start cpu.prof err", err)
		return
	}
	defer pprof.StopCPUProfile()

	// 业务代码
	slowFunc()

	// 获取内存信息
	f1, err := os.Create("mem.prof")
	defer f1.Close()

	if err != nil {
		fmt.Println("create mem.prof err", err)
		return
	}

	if err := pprof.WriteHeapProfile(f1); err != nil {
		fmt.Println("write mem.prof err", err)
		return
	}

	// 获取协程信息
	f2, err := os.Create("goroutine.prof")
	defer f2.Close()
	if err != nil {
		fmt.Println("create goroutine err", err)
		return
	}

	if gProf := pprof.Lookup("goroutine"); gProf != nil {
		fmt.Println("write goroutine err", err)
		return
	} else {
		gProf.WriteTo(f2, 0)
	}

	return

}

func slowFunc() {
	str := "hello, world "
	for i := 0; i <= 5; i++ {
		str += str
	}
}