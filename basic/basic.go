package main

import (
	"fmt"
	"math"
)

var c int
var d bool

func varZero() {
	var a int
	var b string
	fmt.Printf("varZero: %d, %q \n", a, b)
}

func varMultipart() {
	var a, b = 100, "abc"
	fmt.Println("varMultipart:", a, b)
}

func varShort() {
	a, b := 10, "def" // 只能在函数内使用 :=
	fmt.Println("varShort:", a, b)
}

func varAll() {
	fmt.Println("varAll:", c, d)
}

/**
* bool, string

* (u)int(根据系统决定长度), (u)int8, (u)int16, (u)int32, (u)int64, uintptr(指针)

* byte, rune(go语言的char类型)

* float32, float64, complex64, complex128

 */

func transForm() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b))) // go语言只有强制转换，没有隐式转换
	fmt.Println(c)
}

func consts() {
	const (
		filename = "abc.txt"
	)
	fmt.Println(filename)
}

func enums() {
	const (
		js = iota // 自增
		php
		python
	)
	fmt.Println(js, php, python)
}


func main() {
	varZero()
	varMultipart()
	varShort()
	varAll()
	transForm()
	consts()
	enums()
}
