package service

import (
	"github.com/gin-gonic/gin"
	"lite-talk/models"
	"lite-talk/responseutil"
	"net/http"
	"strconv"
)

func Topics(c *gin.Context) {
	p, _ := strconv.Atoi(c.DefaultQuery("p", "1"))
	topics, pages := models.QueryTopics(p, 2)

	var data = make(map[string]interface{})
	data["topics"] = topics
	data["pages"] = pages
	c.JSON(http.StatusOK, responseutil.SuccessWithData(data))
}
