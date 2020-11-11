package main

import (
	"fmt"
	"go-study/array/sort"
)

func main() {
	//ar := [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	//fmt.Println(ar)
	//for index, value := range ar {
	//	for i, v := range value {
	//		ar[i][index] = v
	//	}
	//}
	//
	//fmt.Println(ar)

	var arr = []int{4,78,5,51,5}
	l := len(arr) - 1
	sort.QuickSort(&arr, 0, l)

	//fmt.Println(arr)
	index := sort.BinaryFind(&arr,0, l, 5)
	fmt.Println(index)
}
