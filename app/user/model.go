package user

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/honpery-com/be-api/app/common"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const CollName common.CollName = "users"

type UserId = bson.ObjectId

type UserStatus = string

const (
	Active UserStatus = "active"
	Delete UserStatus = "delete"
)

type UserType = string

const (
	UserTypeAdmin UserType = "admin"
)

type User struct {
	common.BaseModel
	Name     string     `json:"name" bson:"name"`
	Desc     string     `json:"desc" bson:"desc"`
	UserName string     `json:"username" bson:"username"`
	Password string     `json:"password" bson:"password"`
	Site     string     `json:"site" bson:"site"`
	Status   UserStatus `json:"status" bson:"status"`
	Type     UserType   `json:"type" bson:"type"`

	_conn *mgo.Collection
}

func Model(c *gin.Context) *User {
	user := User{}
	user._conn = c.MustGet("mgo_db").(*mgo.Database).C(CollName)
	return &user
}

func (m User) List(conditions interface{}) ([]User, error) {
	result := []User{}
	err := m._conn.Find(conditions).All(&result)
	return result, err
}

func (m User) Detail(user_id string) (User, error) {
	result := User{}
	err := m._conn.FindId(bson.ObjectId(user_id)).One(&result)
	return result, err
}

func (m User) Create(new_user User) (User, error) {
	new_user.Id = bson.NewObjectId()
	new_user.CreateAt = time.Now()
	new_user.UpdateAt = time.Now()
	new_user.Status = Active
	err := m._conn.Insert(new_user)
	return new_user, err
}

func (m User) Update(user_id string, new_user User) (User, error) {
	err := m._conn.UpdateId(bson.ObjectId(user_id), new_user)
	return new_user, err
}

func (m User) Delete(user_id string) {

}
