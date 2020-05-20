package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	user := r.Group("/user")
	{
		user.POST("/register", UserRegister)
		user.POST("/login", UserLogin)
		user.GET("/info/:id", UserInfo)
	}

	block := r.Group("/topics")
	{
		block.Group("/:blockId", TopicsOfBlock)
	}
	return r
}
