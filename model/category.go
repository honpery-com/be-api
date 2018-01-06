package model

import "time"

// CategoryID alias
type CategoryID = string

// CategoryStatus type
type CategoryStatus = int

const (
	// CategoryStatusAvailable = 0
	CategoryStatusAvailable CategoryStatus = iota
)

// Category struct
type Category struct {
	ID       CategoryID     `json:"_id"`
	Name     string         `json:"name"`
	Desc     string         `json:"desc"`
	CreateAt time.Timer     `json:"create_at"`
	UpdateAt time.Timer     `json:"update_at"`
	Status   CategoryStatus `json:"status"`
}
