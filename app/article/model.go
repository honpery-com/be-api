package article

import (
	"time"

	. "github.com/honpery-com/be-api/app/category"
	. "github.com/honpery-com/be-api/app/common"
	. "github.com/honpery-com/be-api/app/tag"
	. "github.com/honpery-com/be-api/app/user"
	"gopkg.in/mgo.v2/bson"
)

type ArticleId bson.ObjectId

const ArticleColl CollName = "articles"

type ArticleStatus string

const (
	Publish ArticleStatus = "publish"
	Draft   ArticleStatus = "draft"
	Delete  ArticleStatus = "delete"
)

type Article struct {
	Id        ArticleId     `json:"_id" bson:"_id,omitempty"`
	Title     string        `json:"title" bson:"title"`
	Desc      string        `json:"desc" bson:"desc"`
	Body      string        `json:"body" bson:"bson"`
	Category  CategoryId    `json:"category" bson:"category"`
	Tags      []TagId       `json:"tags" bson:"tags"`
	Author    UserId        `json:"author" bson:"author"`
	Status    ArticleStatus `json:"status" bson:"status"`
	CreateAt  time.Time     `json:"create_at" bson:"create_at"`
	PublishAt time.Time     `json:"publish_at" bson:"publish_at"`
	UpdateAt  time.Time     `json:"update_at" bson:"update_at"`
}
