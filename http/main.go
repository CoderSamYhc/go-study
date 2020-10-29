package main

import (
	"fmt"
	"net/http"
)

func helloWorld (w http.ResponseWriter, r *http.Request) {
	//_, err := fmt.Fprintf(w, "hello,world")
	//if err != nil {
	//	panic(err)
	//}

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	fmt.Println(r.Form)


}

func main () {

	http.HandleFunc("/hello", helloWorld)
	// 静态文件
	files := http.FileServer(http.Dir("./public")) // 推荐使用绝对路径
	// 将public替换为static
	http.Handle("/static/", http.StripPrefix("/static/", files))
	// 配置监听
	server := http.Server{
		Addr: ":8081",
	}
	// 启动监听
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}


