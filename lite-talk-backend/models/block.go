package models

import "time"

type Block struct {
	Id         int
	Name       string
	TopicCount int
	CreateTime time.Time
	UpdateTime time.Time
}

func QueryBlocks() []Block {
	var blocks []Block
	db.Find(&blocks)
	return blocks
}
