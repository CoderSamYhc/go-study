package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct {
	Name string
	Age int
	Score int
}

type Class []Student

func (c Class) Len() int {
	return len(c)
}

func (c Class) Less(i, j int) bool {
	return c[i].Score > c[j].Score
}

func (c Class) Swap(i, j int) {
	c[i].Score , c[j].Score = c[j].Score , c[i].Score
}
func main() {

	var c Class

	for i := 0; i < 10; i++ {
		stu := Student{
			Name: fmt.Sprintf("姓名：%d", rand.Intn(100)),
			Age: rand.Intn(100),
			Score: rand.Intn(100),
		}

		c = append(c, stu)
	}
	//fmt.Println(c)
	sort.Sort(c)
	for _, v := range c{
		fmt.Println(v)
	}
}
