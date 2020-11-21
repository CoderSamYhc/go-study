package main

import "fmt"


// 传递一个指针地址
func change(p *int) {
	fmt.Println(p, *p)
}

func main() {

	var p int

	fmt.Println(&p) // &取地址符

	change(&p)


}
