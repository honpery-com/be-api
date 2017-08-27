package user

import (
	"time"

	. "../common"
	"gopkg.in/mgo.v2/bson"
)

const UserColl CollName = "users"

type UserId bson.ObjectId

type UserStatus string

const (
	Active UserStatus = "active"
	Delete UserStatus = "delete"
)

type UserType string

const (
	UserTypeAdmin UserType = "admin"
)

type User struct {
	Id       UserId     `json:"_id" bson:"_id,omitempty"`
	Name     string     `json:"name" bson:"name"`
	Desc     string     `json:"desc" bson:"desc"`
	UserName string     `json:"username" bson:"username"`
	Password string     `json:"password" bson:"password"`
	Status   UserStatus `json:"status" bson:"status"`
	Author   UserId     `json:"author" bson:"author"`
	Type     UserType   `json:"type" bson:"type"`
	CreateAt time.Time  `json:"create_at" bson:"create_at"`
	UpdateAt time.Time  `json:"update_at" bson:"update_at"`
}
