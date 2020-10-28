package main

import "fmt"

func main () {
	//var i int = 65536
	//b := int8(i)
	//fmt.Println(b) // 0

	var i int = 1
	i = i << 10 // 1 乘以 2的10次方
	fmt.Println(i)

	var j int = 1024
	j = j >> 10 // 1024 除以 2的10次方
	fmt.Println(j)
}
