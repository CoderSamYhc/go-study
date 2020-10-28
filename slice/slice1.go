package main

import "fmt"

func main() {


	var s1 = make([]int, 5)
	var s2 = make([]int, 5)
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7}



	//s1 = arr[2:6] // [2 3 4 5]
	s1 = arr[2:6:6] // [2 3 4 5] 可以避免 s2读取 arr后面的值

	s2 = s1[1:3] // [2 3 4 5 6]
	s2[0] = 100
	fmt.Printf("s1 = %v \n", s1) // s1 = [2 100 4 5]
	fmt.Printf("s2 = %v \n", s2) // s2 = [100 4 5 6]
	fmt.Printf("arr = %v \n", arr) // [0 1 2 100 4 5 6 7 0 0]
}
