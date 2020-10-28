package user

import "fmt"

type User struct {
	id int
	name string
	age int
}

var data []User

func NewUser() *User {
	return &User{}
}

func (u *User) AddUser(name string, age int) (bool) {
	if len(name) == 0 {
		fmt.Println("名字不能为空")
		return false
	}

	if age < 0 {
		fmt.Println("年龄不能小于0")
		return false
	}
	u.name = name
	u.age = age
	u.id = len(data) + 1
	data = append(data, *u)
	return true
}

func (u User) GetUsers() *[]User {

	return &data
}

func (u User) GetName() string {
	return u.name
}

func (u User) GetAge() int {
	return u.age
}

func (u User) GetId() int {
	return u.id
}

func (u User) GetUser(id int) (bool, *User) {
	if len(data) < id {
		return false, &User{}
	}
	user := data[id - 1]
	return true, &user
}