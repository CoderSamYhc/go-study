package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"go-study/web/middleware"
)

type User struct {
	Username string `from:"username" json:"username" binding:"required"`
	Password string `from:"password" json:"password" binding:"required"`
}

func userRouter(r *gin.Engine) {

	user := r.Group("/user")
	{
		user.POST("/login", middleware.MyFmt(), login)
	}



}

func login (c *gin.Context) {
	var user User
	err := c.ShouldBind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err" : err.Error(),
		})
		return
	}
	fmt.Println("登陆成功..")
	c.JSON(http.StatusOK, gin.H{
		"username" : user.Username,
		"password" : user.Password,
	})
	return
}