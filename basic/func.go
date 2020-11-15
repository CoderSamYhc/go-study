package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (res int, err error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("op / error b not zero, b = %v", b)
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("undefined op %s", op)
	}
}

func exChange(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

// 复合函数
func apply(op func(int, int) int, a, b int) int {
	// 获取某函数名字
	p := reflect.ValueOf(op).Pointer() // 通过反射获取函数的指针
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function %s with args (%d, %d) \n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	res := 0

	for _, v := range numbers {
		res += v
	}
	return res
}
func main() {

	if res, err := eval(100, 0, "x"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	a := 100
	b := 200
	a, b = exChange(a, b)

	fmt.Printf("a = %d, b = %d \n", a, b)

	fmt.Println(apply(pow, 4, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))
}
