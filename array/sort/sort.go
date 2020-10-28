package sort

import "fmt"

func bubbleSort(arr []int) []int {
	var temp int
	// 冒泡
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr) - 1 - i; j++ {
			if arr[j] > arr[j + 1] {
				// 交换
				temp = arr[j]
				arr[j] = arr[j + 1]
				arr[j + 1] = temp
			}
		}
	}
	return arr
}

// 二分查找
func BinaryFind(arr *[]int,lIndex int, rIndex int, findNum int) int {
	if lIndex > rIndex {
		fmt.Println("未找到")
		return -1
	}
	// 计算中间下标
	var middle = (lIndex + rIndex) / 2
	if (*arr)[middle] > findNum { // 在左边
		middle = BinaryFind(arr, lIndex, middle - 1, findNum)
	} else if (*arr)[middle] < findNum {
		middle = BinaryFind(arr, middle + 1, rIndex, findNum)
	}else {
		if middle == 0 || (*arr)[middle-1] != findNum {
			return middle
		} else {
			middle = BinaryFind(arr, lIndex, middle - 1, findNum)
		}
	}
	return middle
}
// 快排
func QuickSort(arr *[]int,  lIndex int, rIndex int) {
	if lIndex >= rIndex {
		return
	}

	var pivot int = (*arr)[rIndex]
	var j = lIndex
	var temp int
	for i := lIndex; i < rIndex; i++ {
		if (*arr)[i] < pivot {
			(*arr)[j], (*arr)[i] = (*arr)[i], (*arr)[j]
			j++
		}
	}
	temp = (*arr)[j]
	(*arr)[j] = pivot
	(*arr)[rIndex] = temp
	QuickSort(arr, lIndex, j - 1)
	QuickSort(arr, j + 1, rIndex)
	return
}

