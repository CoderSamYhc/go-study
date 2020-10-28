package main

import "fmt"

type AInterface interface {
	Say()
}
type BInterface interface {
	Hello()
}
type Integer int

func (i Integer) Say() {
	fmt.Println("integer = ", i)
}
func (i Integer) Hello() {
	fmt.Println("hello")
}
func main() {

	var i Integer = 10

	var ai AInterface = i
	var bi BInterface = i
	ai.Say()
	bi.Hello()
}
