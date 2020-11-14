package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func orderRouter(r *gin.Engine) {

	order := r.Group("/order")
	{
		order.POST("/ms", ms)
	}

}

func ms(c *gin.Context) {
	fmt.Println(c)
}