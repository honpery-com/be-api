package model

import "time"

// TagID alias
type TagID = string

// TagStatus type
type TagStatus = int

// Tag struct
type Tag struct {
	ID       TagID      `json:"_id"`
	Name     string     `json:"name"`
	Desc     string     `json:"desc"`
	CreateAt time.Timer `json:"create_at"`
	UpdateAt time.Timer `json:"update_at"`
	Status   TagStatus  `json:"status"`
}
