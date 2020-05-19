package models

import (
	"time"
)

type Topic struct {
	Id         int
	UserId     int
	BlockId    int
	Title      string
	Content    string
	ClickCount int
	ReplyCount int
	CreateTime time.Time
	UpdateTime time.Time
}

func QueryTopics(page int, pageNum int) (topics []Topic, pages int) {
	topics, count := make([]Topic, 0, pageNum), 0
	db.Debug().Offset((page - 1) * pageNum).Limit(pageNum).Find(&topics)
	db.Table("topic").Count(&count)
	//计算总页数
	pages = count / pageNum
	if count%pageNum != 0 {
		pages++
	}
	return topics, pages
}

func AddTopic(pTopic map[string]interface{}) {
	var userId = pTopic["userId"].(int)

	var topic = Topic{UserId: userId}
}
