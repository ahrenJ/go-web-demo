package routers

import (
	"github.com/gin-gonic/gin"
	"lite-talk/models"
	"lite-talk/responseutil"
	"net/http"
	"strconv"
)

type Topic struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	BlockId int    `json:"blockId" binding:"required"`
}

func PostTopic(c *gin.Context) {
	var body Topic
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseutil.ErrorArgs())
	}

	var topic = models.Topic{Title: body.Title, Content: body.Content, BlockId: body.BlockId}
	models.AddTopic(&topic)
	c.JSON(http.StatusOK, responseutil.SuccessWithMsg("发布主题成功"))
}

func Topics(c *gin.Context) {
	p, _ := strconv.Atoi(c.DefaultQuery("p", "1"))
	topics, pages := models.QueryTopics(p, 10)

	var data = make(map[string]interface{})
	data["topics"] = topics
	data["pages"] = pages
	c.JSON(http.StatusOK, responseutil.SuccessWithData(data))
}
