package category

import (
	"time"

	"github.com/honpery-com/be-api/app/common"

	"gopkg.in/mgo.v2"

	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2/bson"
)

const CollName common.CollName = "categories"

type CategoryId = bson.ObjectId

type CategoryStatus = string

const (
	Active CategoryStatus = "active"
	Delete CategoryStatus = "delete"
)

type Category struct {
	common.BaseModel
	Name   string         `json:"name" bson:"name"`
	Desc   string         `json:"desc" bson:"desc"`
	Status CategoryStatus `json:"status" bson:"status"` // default active

	_conn *mgo.Collection
}

func Model(c *gin.Context) *Category {
	model := Category{}

	model._conn = c.MustGet("mgo_db").(*mgo.Database).C(CollName)

	return &model
}

func (m Category) List(conditions interface{}) ([]Category, error) {
	result := []Category{}
	err := m._conn.Find(conditions).All(&result)
	return result, err
}

func (m Category) Detail(category_id string) (Category, error) {
	result := Category{}
	err := m._conn.FindId(bson.ObjectIdHex(category_id)).One(&result)
	return result, err
}

func (m Category) Create(new_category Category) (Category, error) {
	new_category.Id = bson.NewObjectId()
	new_category.Status = Active
	new_category.CreateAt = time.Now()
	new_category.UpdateAt = time.Now()
	err := m._conn.Insert(new_category)
	return new_category, err
}

func (m Category) Update(category_id string, new_category Category) (Category, error) {
	err := m._conn.UpdateId(category_id, new_category)
	return new_category, err
}

func (m Category) Delete(category_id string) (Category, error) {
	result := Category{}
	result.Status = Delete
	err := m._conn.UpdateId(category_id, result)
	return result, err
}
