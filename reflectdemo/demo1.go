package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age int
}
func (p Person) SayName() {
	fmt.Println(p.Name)
}

func (p Person) SaySomething(str string) {
	fmt.Println(str)
}

func (p Person) SumNum(num ...int) int {
	var rs int
	for _, v := range num {
		rs += v
	}

	return rs
}

func main() {

	var stu = Person{"Sam",19}
	
	//fmt.Println(reflect.ValueOf(stu))
	//fmt.Println(reflect.ValueOf(stu).Type()) // 等同于 reflect.TypeOf(stu)
	//fmt.Println(reflect.TypeOf(stu)) // main.Person
	//fmt.Println(reflect.TypeOf(stu).Name()) // Person 类型名 type
	//fmt.Println(reflect.TypeOf(stu).Kind()) // struct 类别名 kind
	fv := reflect.ValueOf(stu)
	// 调用无参数的方法
	f := fv.MethodByName("SayName")
	rr := f.Call(nil)
	fmt.Println(rr)
	// 调用有参数的方法

	f2 := fv.MethodByName("SaySomething")
	p := []reflect.Value{reflect.ValueOf("hello")}
	r := f2.Call(p)
	fmt.Println(r)
	f3 := fv.MethodByName("SumNum")
	p2 := []reflect.Value{
		reflect.ValueOf(20),
		reflect.ValueOf(10),
	}

	rs := f3.Call(p2)
	fmt.Printf("%T %v %T\n", rs, reflect.ValueOf(rs), rs[0])
	fmt.Println(rs[0])
	//var num = 100
	//
	//rNum := reflect.ValueOf(num) // 值传递
	//fmt.Println(rNum.Int())
	//fmt.Println(rNum.CanSet()) // 是否可修改
	//fmt.Println(rNum) // 100
	//fmt.Println("----------")
	//
	//rNum2 := reflect.ValueOf(&num) // 引用传递
	//newNum := rNum2.Elem() // 获取反射对象
	//fmt.Println(newNum.CanSet()) // 是否可修改
	//fmt.Println(newNum.Int())
	//newNum.SetInt(200)
	//fmt.Println(newNum.Int())
	//fmt.Println(rNum2) // 地址


}
