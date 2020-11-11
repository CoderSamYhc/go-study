package main

import (
	"fmt"
	"html/template"
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
	fmt.Println("path", r.URL.Path)
	fmt.Println(r.Form)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./public/html/index.html")
	t.Execute(w, "hello golang")
}
func handleFunc() {
	http.HandleFunc("/hello/", helloWorld)
	http.HandleFunc("/", index)
}

func muxFunc() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello/", helloWorld) // /hello 会响应helloWorld方法，/hello/会响应index方法
	mux.HandleFunc("/", before(index))
	return mux;
}

func before(handle http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before.....")
		handle(w, r)
	}
}

func main () {

	//handleFunc();
	mux := muxFunc();
	// 静态文件
	files := http.FileServer(http.Dir("./public")) // 推荐使用绝对路径
	// 将public替换为static
	http.Handle("/static/", http.StripPrefix("/static/", files))
	// 配置监听
	server := http.Server{
		Addr: ":8081",
		Handler: mux,
	}
	// 启动监听
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}


