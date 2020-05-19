package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)

	r := gin.Default()
	r.GET("/hi", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "I'm gin",
		})
	})

	user := r.Group("/user")
	{
		user.POST("/register", UserRegister)
		user.POST("/login", UserLogin)
		user.GET("/info/:id", UserInfo)
	}

	topic := r.Group("/topic")
	{
		topic.GET("/all", Topics)
	}
	return r
}
