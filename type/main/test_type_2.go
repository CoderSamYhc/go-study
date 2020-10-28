package main

import (
	"bytes"
)

func main() {

	//type Person struct {
	//	name string
	//	age int
	//}
	//
	//m := map[string]Person{"key" : {"sam", 18}}
	////for i, v := range m {
	////	fmt.Println(i, v)
	////}
	//m["no"] = Person{"alpha", 12}
	//delete(m, "key")
	//fmt.Println(m)
	//p, j := m["key"] // 返回两个值，1：当前key的值，2：key是否存在的bool值
	//fmt.Println(p) // {sam 18} key 不存在返回默认值 { 0}
	//fmt.Println(j) // true key 不存在返回false
	//fmt.Println(m) // map[key:{sam 18}]
	//fmt.Println(m["key"].name) // sam
	//
	//type Student struct {
	//	Person
	//	className string
	//}
	//
	//stu := Student{
	//	Person{"Sam", 19},
	//	"on_1",
	//}
	//
	//fmt.Println(stu.age, stu.name) // 19 Sam
	//var stu = Person{"Sam", 18}
	//stu2 := Person{age:19}
	//fmt.Println(stu2.age) // 19
	//fmt.Println(stu) // {Sam 18}

	// 申请构造体
	//stu := new(Person) // 返回的是指针类型
	//stu.name = "Sam"
	//fmt.Printf("%T \n", stu)



	// 切片
	//var s = []byte{1, 2, 3, 4}
	//test(s)

	//i := 2
	//fmt.Println(s[:i]) // [1 2]
	//fmt.Println(s[i+1:]) // [4]
	//s = append(s[:i], s[i+1:]...)
	//fmt.Println(s) // [1 2 4]
	//s = append(s, 3, 1)
	//fmt.Println(s) // 2
	//s1 := make([]byte, 2)
	//num := copy(s1, s)
	//fmt.Println(num, s1)
	//s := make([]byte, 5)
	//for i, _ := range s{
	//	s[i] = byte(i) // 由于i是int类型 这里转成byte
	//}
	//var s1 = []byte{}
	//var num = copy(s, s1)
	//fmt.Println(num)

	// 指针
	//var a = "abc"
	//var p *string = &a
	//*p = "ddd" // 修改a的值 等价于 a = "ddd"
	//fmt.Println(a)
	//fmt.Println(p) // 0xc000010200
	//fmt.Println(*p)
	//var p *bool
	//p = new(bool)
	//fmt.Println(p)
	//fmt.Println(*p)

	// 数组
	// var arr1 = [10]int{1} // 定义并初始化
	//arr2 := [...]int{1,2,3} // 自动计算长度
	//arr3 := [...]int{2:9,3:3} // 选择性初始化

	//var a = arr1[5:] // 从索引4开始，获取后面的元素
	//var b = arr1[:5] // 到索引4结束，获取前面的元素
	//var c = arr1[:] // 获取全部的元素
	//fmt.Printf("%T \n", a)
	//fmt.Printf("%T \n", arr1)
	//fmt.Println(a, b, c, len(a))

	//for i := 0; i < len(arr1); i++ {
	//	fmt.Println(arr1[i])
	//}
	//
	//for i, v := range arr1 {
	//	fmt.Println(i, v)
	//}

	// 集合
	//m := map[string]int{"a":1, "b":2}
	//fmt.Println(m["a"])

	//m := make(map[string]int)
	//m["a"] = 1
	//fmt.Println(m)

	//var m sync.Map
	//m.Store("key", 10) // 增加元素
	//fmt.Println(m.Load("key")) // 10 true 获取元素 返回两个值，key对应的值， key是否存在的值
	//// 遍历
	//m.Range(func (k, v interface{}) bool{
	//	fmt.Println(k, v)
	//	return true
	//})

	//fmt.Println(joinString("hello", " world", "!!!!", "~~~~"))

	//s := "hello"
	//func (str string) {
	//	fmt.Println(s)
	//}(s)
	//i := 10
	//x := 1
	//max, min := func(a, b int) (max, min int) {
	//	if a > b {
	//		max = a
	//		min = b
	//	} else {
	//		max = b
	//		min = a
	//	}
	//	return
	//}(i, x)
	//
	//fmt.Println(max, min)
}

func joinString(s ...string) string {
	var buf bytes.Buffer

	for _, v := range s {
		buf.WriteString(v)
	}
	return buf.String()
}

//func test(s []byte) {
//	fmt.Println("test---", s)
//}