package tag

import (
	"time"

	. "../common"
	"github.com/honpery-com/be-api/app/user"
	"gopkg.in/mgo.v2/bson"
)

const TagColl CollName = "tags"

type TagId bson.ObjectId

type TagStatus string

const (
	Active TagStatus = "active"
	Delete TagStatus = "delete"
)

type Tag struct {
	Id       TagId       `json:"_id" bson:"_id,omitempty"`
	Name     string      `json:"name" bson:"name"`
	Desc     string      `json:"desc" bson:"desc"`
	Status   TagStatus   `json:"status" bson:"status"`
	Author   user.UserId `json:"author" bson:"author"`
	CreateAt time.Time   `json:"create_at" bson:"create_at"`
	UpdateAt time.Time   `json:"update_at" bson:"update_at"`
}
