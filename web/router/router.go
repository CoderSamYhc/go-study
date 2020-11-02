package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter () *gin.Engine {
	r := gin.Default()

	// 路由
	userRouter(r)

	return r
}
