package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type Student struct {
	Name string `json:"my_name"` // 反射机制，{"my_name":"Sam","my_age":18,"my_height":182}
	Age int `json:"my_age"`
	Height int `json:"my_height"`
}

func main() {
	//testStruct()
	//testMap()
	//testSlice()
	//unmarshalStruct()
	//unmarshalMap()
	//unmarshalSlice()

	jsonStr := `{"Name":"Lisi","Age":6,"Parents":["Lisan","WW"]}`
	jsonBytes := []byte(jsonStr)

	var i interface{}
	json.Unmarshal(jsonBytes, &i)
	fmt.Println(i) // map[Age:6 Name:Lisi Parents:[Lisan WW]]

	m := i.(map[string]interface{})
	for k, v := range m {
		switch r := v.(type) {
		case string:
			fmt.Println(k, " string ", r)
		case int:
			fmt.Println(k, " int ", r)
		case []interface{}:
			fmt.Println(k, " array ", r)
		default:
			fmt.Println("cannot be type")
		}
	}

	fmt.Println(m["Age"])
}

func testStruct() {
	// 构造体转json
	stu := Student{"Sam", 18, 182}
	data, err := json.Marshal(&stu)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func testMap() {
	// 集合转json
	m := make(map[string]string)
	m["name"] = "Sam"
	m["age"] = "18"
	data, err := json.Marshal(m)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}

func testSlice() {
	// 切片转json
	s := make([]int, 4)
	for i, _  := range s {
		s[i] = rand.Intn(100)
	}
	data, err := json.Marshal(s)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}

func unmarshalStruct() {

	// 需要注意结构体是否反射
	str := "{\"my_name\":\"Sam\",\"my_age\":18,\"my_height\":182}"

	var p Student
	err := json.Unmarshal([]byte(str), &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p, p.Age)
}

func unmarshalMap() {
	m := make(map[string]string)
	str := "{\"age\":\"18\",\"name\":\"Sam\"}"
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v \n", m)
}

func unmarshalSlice() {
	s := make([]string ,4)
	str := "[\"sam\",\"18\"]"
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s[0])
}