package main

import (
	"fmt"
)

type Log interface {
	Writer()
	Reader()
}

type LogPrint interface {
	Log
	Print()
}

type Database struct {
	Name string
}

func (db Database) Writer() {
	fmt.Println("log writer in database")
}

func (db Database) Reader() {
	fmt.Println("read log to database")
}

func (db Database) Print() {
	fmt.Println("print log to database")
}

func (db Database) Delete() {
	fmt.Println("log delete in database")
}

type File struct {
	Name string
}

func (f File) Writer() {
	fmt.Println("log writer in file")
}

func (f File) Reader() {
	fmt.Println("read log to file")
}

func (f File) Print() {
	fmt.Println("print log to file")
}

type Program struct {

}

// 接收一个接口类型
// 就是实现了 usb 接口的所有方法的 类
func (p Program) Working(log Log) {
	log.Writer()

	// 类型断言，由于接口是一般类型，不知道具体类型，如果需要转成具体类型，要使用类型断言
	if y, e := log.(Database); e {
		y.Delete()
	}

	log.Reader()
}

func main() {
	p := Program{}
	db := Database{}
	file := File{}
	//
	p.Working(db)
	p.Working(file)
	//
	//// 直接调用
	////var log Log = &db
	//var logp LogPrint = &db
	//logp.Writer()
	//logp.Reader()
	//logp.Print()
	//
	////log = &file
	//logp = &file
	//logp.Writer()
	//logp.Reader()
	//logp.Print()

	//var log [2]Log
	//log[0] = Database{"database"}
	//log[1] = File{"file"}
	//log[0].Writer()

	var i interface{}
	var f float64 = 3.42
	i = f
	// 类型断言
	if y, e := i.(float32); e {
		fmt.Printf("%T %v", y, y)
	} else {
		fmt.Println(e)
	}
}

