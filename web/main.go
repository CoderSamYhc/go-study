package main

import (
	gin "github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `from:"username" json:"username" binding:"required"`
	Password string `from:"password" json:"password" binding:"required"`
}

func login (c *gin.Context) {
	var user User
	err := c.ShouldBind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err" : err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"username" : user.Username,
		"password" : user.Password,
	})
}

func main () {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		username := c.Query("username")

		c.JSON(200, gin.H{
			"msg" : "hello " + username,
		})
	})

	r.Run(":8081")
}
