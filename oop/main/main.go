package main

import (
	"studented/oop/person"
)
func main() {
	//p := person.NewPerson("Sam", 18)
	//fmt.Println(p)
	//p.SetAge(10)
	//p.GetAge()
	//p.Say()
	//p.Run()
	//p.GetAge()
	//fmt.Println(p)
	//var n1 int
	//fmt.Scanln(&n1)
	//p.MethodCal(n1)

	a := person.NewStudent("Visige", 1)
	a.Working(a)
}