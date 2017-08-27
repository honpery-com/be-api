package category

import (
	"time"

	. "github.com/honpery-com/be-api/app/common"
	"github.com/honpery-com/be-api/app/user"
	"gopkg.in/mgo.v2/bson"
)

const CategoryColl CollName = "categories"

type CategoryId bson.ObjectId

type CategoryStatus string

const (
	Active CategoryStatus = "active"
	Delete CategoryStatus = "delete"
)

type Category struct {
	Id       CategoryId     `json:"_id,omitempty" bson:"_id"`
	Name     string         `json:"name" bson:"name"`
	Desc     string         `json:"desc" bson:"desc"`
	status   CategoryStatus `json:"status" bson:"status"`
	Author   user.UserId    `json:"author" bson:"author"`
	CraeteAt time.Time      `json:"create_at" bson:"create_at"`
	UpdateAt time.Time      `json:"update_at" bson:"update_at"`
}

func (m Category) List() {

}

func (m Category) Detail() {

}

func (m Category) Create() {

}

func (m Category) Update() {

}

func (m Category) Delete() {

}
