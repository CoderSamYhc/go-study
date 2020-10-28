package main

import (
	"fmt"
	"reflect"
)

func reflectTest(p interface{}) {

	rType := reflect.TypeOf(p)
	fmt.Println("type=", rType) // type float64
	fmt.Println("kind=", rType.Kind()) // kind float64
	rValue := reflect.ValueOf(p).Interface()
	fmt.Printf("转Interface %T \n", rValue)

	num := rValue.(float64)
	fmt.Printf("转float %T \n", num)
}


func main() {
	var num  = 1.2
	reflectTest(num)
}
