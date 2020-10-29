package main

import (
	. "fmt" // . 可省略报名调用其中方法
	_ "go-study/bao/demo" // _ 会执行该包的init方法，如果存在的话
)

func init () {
	Println("init....")
}

func main () {
	Println("hello,world")
}

