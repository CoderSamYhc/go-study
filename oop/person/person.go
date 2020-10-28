package person

import (
	"fmt"
)

type Action interface {
	Run()
	Say()
}

type Person struct {
	name string
	age int
}


func NewPerson(name string, age int) *Person {
	return &Person{name, age}
}
func (p Person) Say() {
	fmt.Println(p.name, "say: hello")
}
func (p *Person) Run() {
	fmt.Println(p.name, "running!")
	p.age += 1
}

func (p *Person) SetAge (age int) {
	p.age = age
}

func (p Person) GetAge() {
	fmt.Println(p.age)
}

func (p *Person) String() string {
	return fmt.Sprintf("name=%v, age=%v", p.name, p.age)
}

//func (p Person) MethodCal(n1 int) {
//
//	for i := 1; i <= n1; i++ {
//		for j := 1; j <= i; j++ {
//			fmt.Printf("%v * %v = %v \t", i, j, i*j)
//		}
//		fmt.Println()
//	}
//}

type Student struct {
	Person
}
func NewStudent (name string, age int) *Student {
	return &Student{Person{name, age}}
}

func (s Student) Working(action Action) {
	action.Run()
	action.Say()
}
func (s Student) Run() {
	fmt.Println("Running.....")
}

func (s Student) Say() {
	fmt.Println("Say Same thing...")
}
