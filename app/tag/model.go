package tag

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/honpery-com/be-api/app/common"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const CollName common.CollName = "tags"

type TagId bson.ObjectId

type TagStatus string

const (
	Active TagStatus = "active"
	Delete TagStatus = "delete"
)

type Tag struct {
	common.BaseModel
	Name   string    `json:"name" bson:"name"`
	Desc   string    `json:"desc" bson:"desc"`
	Status TagStatus `json:"status" bson:"status"`

	_conn *mgo.Collection
}

func Model(c *gin.Context) *Tag {
	tag := Tag{}

	tag._conn = c.MustGet("mgo_db").(*mgo.Database).C(CollName)

	return &tag
}

func (m Tag) List(conditions interface{}) ([]Tag, error) {
	result := []Tag{}
	err := m._conn.Find(conditions).All(&result)
	return result, err
}

func (m Tag) Detail(tag_id string) (Tag, error) {
	result := Tag{}
	err := m._conn.FindId(bson.ObjectId(tag_id)).One(&result)
	return result, err
}

func (m Tag) Create(new_tag Tag) (Tag, error) {
	new_tag.Id = bson.NewObjectId()
	new_tag.Status = Active
	new_tag.CreateAt = time.Now()
	new_tag.UpdateAt = time.Now()
	err := m._conn.Insert(new_tag)
	return new_tag, err
}

func (m Tag) Update(tag_id string, new_tag Tag) (Tag, error) {
	err := m._conn.UpdateId(tag_id, new_tag)
	return new_tag, err
}

func (m Tag) Delete(tag_id string) (Tag, error) {
	result := Tag{}
	result.Status = Delete
	err := m._conn.UpdateId(tag_id, result)
	return result, err
}
