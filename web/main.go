package main

import (
	"go-study/web/router"
)

func main () {
	r := router.InitRouter()
	_ = r.Run(":8081")
}
