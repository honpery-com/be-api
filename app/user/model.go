package user

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const UserColl = "users"

type UserId bson.ObjectId

type UserStatus string

const (
	UserStatusActive UserStatus = "active"
	UserStatusDelete UserStatus = "delete"
)

type UserType string

const (
	UserTypeAdmin UserType = "admin"
)

type User struct {
	Id       UserId     `json:"_id" bson:"_id"`
	Name     string     `json:"name" bson:"name"`
	Desc     string     `json:"desc" bson:"desc"`
	UserName string     `json:"username" bson:"username"`
	Password string     `json:"password" bson:"password"`
	Status   UserStatus `json:"status" bson:"status"`
	Author   UserId     `json:"author" bson:"author"`
	Type     UserType   `json:"type" bson:"type"`
	CreateAt time.Timer `json:"create_at" bson:"create_at"`
	UpdateAt time.Timer `json:"update_at" bson:"update_at"`
}
