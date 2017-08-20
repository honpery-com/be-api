package tag

const TagColl = "tags"

type Tag struct {
	Name string `json:"name" bson:"name"`
	Desc string `json:"desc" bson:"desc"`
}
