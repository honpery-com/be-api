package article

import (
	"time"

	"github.com/honpery-com/be-api/app/category"
	"github.com/honpery-com/be-api/app/tag"
	"github.com/honpery-com/be-api/app/user"
	"gopkg.in/mgo.v2/bson"
)

type ArticleId bson.ObjectId

const ArticleColl = "articles"

type ArticleStatus string

const (
	ArticleStatusPublish ArticleStatus = "publish"
	ArticleStatusDraft   ArticleStatus = "draft"
	ArticleStatusDelete  ArticleStatus = "delete"
)

type Article struct {
	Id        ArticleId           `json:"_id" bson:"_id"`
	Title     string              `json:"title" bson:"title"`
	Desc      string              `json:"desc" bson:"desc"`
	Body      string              `json:"body" bson:"bson"`
	Category  category.CategoryId `json:"category" bson:"category"`
	Tags      []tag.TagId         `json:"tags" bson:"tags"`
	Author    user.UserId         `json:"author" bson:"author"`
	Status    ArticleStatus       `json:"status" bson:"status"`
	CreateAt  time.Timer          `json:"create_at" bson:"create_at"`
	PublishAt time.Timer          `json:"publish_at" bson:"publish_at"`
	UpdateAt  time.Timer          `json:"update_at" bson:"update_at"`
}
