package tags

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type TagId = sql.NullInt64

type Tag struct {
	Id          TagId          `sql: "index, not null" json: "id"`
	Name        sql.NullString `sql: "type: varchar(100)" json: "name"`
	Description sql.NullString `sql: "type: varchar(200)" json: "name"`
	CreateAt    sql.NullString `sql: "type: timestamp, not null" json: "create_at"`
	UpdateAt    sql.NullString `sql: "type: timestamp, not null" json: "update_at"`

	db *sql.DB
}

func TagModel(ctx *gin.Context) *Tag {
	tag := &Tag{}
	tag.db = ctx.MustGet("DB").(*sql.DB)
	return tag
}

func (m *Tag) List(conditions interface{}) (*[]Tag, error) {
	tags := &[]Tag{}
	query, err := m.db.Query("select * from tags")

	if err != nil {
		return tags, err
	}

	err = query.Scan(tags)
	return tags, err
}

func (m *Tag) Detail(id TagId) (*Tag, error) {
	tag := &Tag{}
	err := m.db.QueryRow("select * from tags where id = $1", id).Scan(tag)
	return tag, err
}

// func (m *Tag) Create(newTag Tag) (*Tag, error) {

// }

// func (m *Tag) Update(id TagId, newTag Tag) (*Tag, error) {

// }

// func (m *Tag) Delete(id TagId) (*Tag, error) {

// }
