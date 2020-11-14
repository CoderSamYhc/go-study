package main

import (
	"go-study/ms/router"
)

func main() {
	r := router.InitRouter()

	_ = r.Run(":8081")
}
