package tag

import (
	"time"

	"github.com/honpery-com/be-api/app/user"
	"gopkg.in/mgo.v2/bson"
)

const TagCell = "tags"

type TagId bson.ObjectId

type TagStatus string

const (
	TagStatusActive TagStatus = "active"
	TagStatusDelete TagStatus = "delete"
)

type Tag struct {
	Id       TagId       `json:"_id" bson:"_id"`
	Name     string      `json:"name" bson:"name"`
	Desc     string      `json:"desc" bson:"desc"`
	Status   TagStatus   `json:"status" bson:"status"`
	Author   user.UserId `json:"author" bson:"author"`
	CreateAt time.Timer  `json:"create_at" bson:"create_at"`
	UpdateAt time.Timer  `json:"update_at" bson:"update_at"`
}
