package routers

import (
	"github.com/gin-gonic/gin"
	"lite-talk/service"
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
	//用户模块api
	user := r.Group("/user")
	{
		user.POST("/register", service.UserRegister)
		user.POST("/login", service.UserLogin)
		user.GET("/info/:id", service.UserInfo)
	}
	//主题模块api
	topic := r.Group("/topic")
	{
		topic.GET("/all", service.Topics)
	}
	return r
}
