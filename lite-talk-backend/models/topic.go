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
	db.Offset((page - 1) * pageNum).Limit(pageNum).Find(&topics)
	db.Table("topic").Count(&count)

	pages = count / pageNum
	if count%pageNum != 0 {
		pages++
	}
	return topics, pages
}

func QueryTopicsByBlockId(blockId int, page int, pageNum int) (topics []Topic, pages int) {
	topics, count := make([]Topic, 0, pageNum), 0
	db.Where("block_id = ?", blockId).Offset((page - 1) * pageNum).Limit(pageNum).Find(&topics)
	db.Table("topic").Count(&count)

	pages = count / pageNum
	if count%pageNum != 0 {
		pages++
	}
	return topics, pages
}

func AddTopic(topic *Topic) {
	topic.CreateTime = time.Now()
	topic.UpdateTime = time.Now()
	db.Create(&topic)
}
