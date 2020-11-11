package main

import (
	"fmt"
	"os"
)

func main() {
	// file 是文件的地址
	//var file, err = os.Open("content.text")
	//fmt.Println(file)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer file.Close()  // 关闭文件句柄
	//if err != nil {
	//	fmt.Println(err)
	//}
	//// 缓冲读取
	//reader := bufio.NewReader(file)
	//for {
	//	str, err := reader.ReadString('\n') // 读到换行就结束
	//	if err == io.EOF { // io.EOF表示文件末尾
	//		break
	//	}
	//	fmt.Printf("%v", str)
	//}

	// 适合小文件
	//content, errs := ioutil.ReadFile("content.text")
	//
	//if errs != nil {
	//	fmt.Println(errs)
	//}
	//fmt.Println(string(content))

	// 写文件
	//file, err := os.OpenFile("hello.log", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer file.Close()
	//
	//str := "world  \n"
	//writer := bufio.NewWriter(file) // 写入缓冲
	//for i := 1; i <= 5; i++ {
	//	//n1, _ := file.WriteString(str) // 直接写入文件
	//	n1, _ := writer.WriteString(str)
	//	fmt.Println(n1)
	//}
	//writer.Flush() // 缓冲写入文件

	//file, err := os.OpenFile("content.text", os.O_RDONLY, 0666)
	//file2, err2 := os.OpenFile("content2.text", os.O_WRONLY | os.O_CREATE, 0666)
	//
	//defer file.Close()
	//defer file2.Close()
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	//reader := bufio.NewReader(file)
	//writer := bufio.NewWriter(file2)
	//
	//io.Copy(writer, reader)

	//for {
	//	r1, ferr := reader.ReadString('\n')
	//	if ferr == io.EOF {
	//		break
	//	}
	//	rs, ferr2 := file2.WriteString(r1)
	//	fmt.Println(rs, ferr2)
	//}

	// err == nil 表示文件存在
	//_, err := os.Stat("content2.text")
	//rs := os.IsNotExist(err)
	//if  !rs {
	//	fmt.Println("存在")
	//}

	p, err := os.Executable()
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(p)

}