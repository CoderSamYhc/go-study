package main

import (
	"errors"
	"fmt"
)

func test(i int) (err error) {
	if i == 10 {
		return nil
	} else {
		return errors.New("创建一个错误")
	}
}

func test2() {

	defer func() {
		errs := recover()
		if errs != nil {
			fmt.Println(errs)
		}
	}()
	err := test(1)
	if err != nil {
		panic(err) // recover()可捕获该异常
	}
	fmt.Println("ok") // 不会被执行
}
func main() {
	// 当执行到defer时，不会立即执行，会将后面的语句压入 defer栈
	// 程序执行完毕时，会以先入后出的顺序执行
	//defer fmt.Println("defer_1")
	//defer fmt.Println("defer_2")
	//fmt.Println("3")


	test2()
	fmt.Println(11) // 会执行
}