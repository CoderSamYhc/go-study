package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func MyFmt() gin.HandlerFunc {
	return func(c *gin.Context) {
		host := c.Request.Host
		fmt.Printf("before:%s\n", host)
		c.Next()
		fmt.Println("next....")
	}
}

