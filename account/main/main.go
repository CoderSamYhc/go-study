package main

import (
	"fmt"
	"studented/account/user"
)

func main() {

	var et bool
	for {
		fmt.Println("-------账号信息录入--------")
		fmt.Printf("\t1.录入信息\n")
		fmt.Printf("\t2.信息列表\n")
		fmt.Printf("\t3.查看明细\n")
		fmt.Printf("\t4.退出\n")
		fmt.Printf("\t请输入1-4\n")

		var scanNum int
		fmt.Scanln(&scanNum)

		u := user.NewUser()
		switch scanNum {
		case 1:
			var name string
			var age int
			fmt.Println("请输入姓名：")
			fmt.Scanln(&name)
			fmt.Println("请输入年龄：")
			fmt.Scanln(&age)
			u.AddUser(name, age)
		case 2:
			data := u.GetUsers()
			if len(*data) == 0 {
				fmt.Println("暂无数据")
				break
			}
			fmt.Println("-------账号信息列表--------")
			for _, v := range *data {
				fmt.Printf("ID：%v, 姓名：%v, 年龄：%v \n", v.GetId(),  v.GetName(), v.GetAge())
			}
			fmt.Printf("-------账号信息列表-------- \n")
		case 3:
			var id int
			fmt.Scanln(&id)
			check, data := u.GetUser(id)
			if !check {
				fmt.Println("暂无数据")
				break
			}
			fmt.Println("-------账号信息详情--------")
			fmt.Printf("ID：%v, 姓名：%v, 年龄：%v \n", data.GetId(), data.GetName(), data.GetAge())
			fmt.Printf("-------账号信息详情-------- \n\n")
		case 4:
			et = true
		default:
		}

		if et {
			break
		}
	}

}