package category

import (
	"time"

	"github.com/honpery-com/be-api/app/user"
	"gopkg.in/mgo.v2/bson"
)

const CategoryCell = "categories"

type CategoryId bson.ObjectId

type CategoryStatus string

const (
	CategoryStatusActive CategoryStatus = "active"
	CategoryStatusDelete CategoryStatus = "delete"
)

type Category struct {
	Id       CategoryId     `json:"_id" bson:"_id"`
	Name     string         `json:"name" bson:"name"`
	Desc     string         `json:"desc" bson:"desc"`
	status   CategoryStatus `json:status bson:status`
	Author   user.UserId    `json:"author" bson:"author"`
	CraeteAt time.Timer     `json:"create_at" bson:"create_at"`
	UpdateAt time.Timer     `json:"update_at" bson:"update_at"`
}
