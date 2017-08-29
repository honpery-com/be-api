package common

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/honpery-com/be-api/app/user"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BaseModel struct {
	Id       bson.ObjectId `json:"_id" bson:"_id"`
	Author   user.UserId   `json:"author" bson:"author"`
	CreateAt time.Time     `json:"create_at" bson:"create_at"`
	UpdateAt time.Time     `json:"update_at" bson:"update_at"`
}

type CollName = string

type Model struct {
	conn   *mgo.Collection
	schema interface{}
}

func XModel(c *gin.Context, name string) *Model {
	model := Model{}

	model.conn = c.MustGet("mgo_db").(*mgo.Database).C(name)

	return &model
}

func (m Model) List(coditions interface{}) (interface{}, error) {
	var result []interface{}
	err := m.conn.Find(coditions).All(&result)
	return result, err
}

func (m Model) Detail(id string) (interface{}, error) {
	var result interface{}
	err := m.conn.FindId(bson.ObjectIdHex(id)).One(&result)
	return result, err
}

func (m Model) Create(newData interface{}) {

	// newData.Id = bson.NewObjectId()
	// newData.CraeteAt = time.Now()
	// newData.UpdateAt = time.Now()
}

func (m Model) Update() {

}

func (m Model) Delete() {

}
