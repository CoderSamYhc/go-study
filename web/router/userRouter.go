package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `from:"username" json:"username" binding:"required"`
	Password string `from:"password" json:"password" binding:"required"`
}

func userRouter(r *gin.Engine) {

	user := r.Group("/user")
	{
		user.POST("/login", login)
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
	c.JSON(http.StatusOK, gin.H{
		"username" : user.Username,
		"password" : user.Password,
	})
	return
}