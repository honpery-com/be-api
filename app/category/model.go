package category

type CategoryID bson.ObjectID

type CategoryStatus int

const (
	Active CategoryStatus = iota
	Delete
)

// type Category struct {
// 	Id       CategoryID `json:"_id" bson:"_id"`
// 	Name     string     `json:"name" bson:"name"`
// 	CraeteAt time.time  `json:"create_at" bson:"create_at"`
// 	UpdateAt time.time  `json:"update_at" bson:"update_at"`
// }
