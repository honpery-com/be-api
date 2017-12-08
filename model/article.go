package model

import "time"

// ArticleStatus is int value
type ArticleStatus = int

const (
	// ArticleStatusPublished = 0
	ArticleStatusPublished ArticleStatus = iota
)

// Article entity struct
type Article struct {
	ID       string        `json:"_id"`
	Title    string        `json:"title"`
	Desc     string        `json:"desc"`
	Body     string        `json:"body"`
	CreateAt time.Timer    `json:"create_at"`
	UpdateAt time.Timer    `json:"update_at"`
	Status   ArticleStatus `json:"status"`
}
