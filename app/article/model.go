package article

import (
	"time"

	"gopkg.in/mgo.v2"

	"github.com/gin-gonic/gin"
	"github.com/honpery-com/be-api/app/category"
	"github.com/honpery-com/be-api/app/common"
	"github.com/honpery-com/be-api/app/tag"
	"gopkg.in/mgo.v2/bson"
)

type ArticleId bson.ObjectId

const CollName common.CollName = "articles"

type ArticleStatus string

const (
	Publish ArticleStatus = "publish"
	Draft   ArticleStatus = "draft"
	Delete  ArticleStatus = "delete"
)

type Article struct {
	common.BaseModel
	Title     string              `json:"title" bson:"title"`
	Desc      string              `json:"desc" bson:"desc"`
	Body      string              `json:"body" bson:"bson"`
	Category  category.CategoryId `json:"category" bson:"category"`
	Tags      []tag.TagId         `json:"tags" bson:"tags"`
	Status    ArticleStatus       `json:"status" bson:"status"`
	PublishAt time.Time           `json:"publish_at" bson:"publish_at"`

	_conn *mgo.Collection
}

func Model(c *gin.Context) *Article {
	article := Article{}

	article._conn = c.MustGet("mgo_db").(*mgo.Database).C(CollName)

	return &article
}

func (m Article) List(conditions interface{}) ([]Article, error) {
	result := []Article{}
	err := m._conn.Find(conditions).all(&result)
	return result, err
}

func (m Article) Detail(article_id string) (Article, error) {
	result := Article{}
	err := m._conn.FindId(bson.ObjectId(article_id)).one(&result)
	return result, err
}

func (m Article) Create(new_article Article) (Article, error) {
	new_article.Id = bson.NewObjectId()
	new_article.CreateAt = time.Now()
	new_article.UpdateAt = time.Now()
	new_article.Status = ArticleStatus
	err := m._conn.Insert(new_article)
	return new_article, err
}

func (m Article) Update(article_id string, new_article Aricle) (Article, error) {
	err := m._conn.UpdateId(article_id, new_article)
	return new_article, err
}

func (m Article) Delete(article_id string) (Article, error) {
	result := Article{}
	result.Status = Delete
	err := m._conn.UpdateId(article_id, result)
	return result, err
}
