package common

import (
	"fmt"

	. "github.com/honpery-com/be-api/config"
	mgo "gopkg.in/mgo.v2"
)

type CollName string

func Connect(coll CollName) (error, *mgo.Session, *mgo.Collection) {
	mongo_url := fmt.Sprintf("%s:%d", DBHost, DBPort)
	session, err := mgo.Dial(mongo_url)

	if err != nil {
		return err, nil, nil
	}

	session.SetMode(mgo.Monotonic, true)
	return nil, session, session.DB(DBName).C(string(coll))
}

type xmodel struct {
	Name CollName
}

func XModel(name CollName) xmodel {
	return xmodel{Name: name}
}

func (m xmodel) List(conditions interface{}) (error, interface{}) {
	err, session, coll := Connect(m.Name)
	defer session.Close()

	if err != nil {
		return err, nil
	}

	var _result interface{}

	err = coll.Find(conditions).All(_result)

	return err, _result
}

func (m xmodel) Detail(id string) (error, interface{}) {
	err, session, coll := Connect(m.Name)
	defer session.Close()

	if err != nil {
		return err, nil
	}

	var _result interface{}

	err = coll.FindId(id).One(_result)

	return err, _result
}

func (m xmodel) Create(doc interface{}) (error, interface{}) {
	err, session, coll := Connect(m.Name)
	defer session.Close()

	if err != nil {
		return err, nil
	}

	err = coll.Insert(doc)

	if err != nil {
		return err, nil
	}

	return nil, doc
}

func (m xmodel) Update(id string, doc interface{}) (error, interface{}) {
	err, session, coll := Connect(m.Name)
	defer session.Close()

	if err != nil {
		return err, nil
	}

	err = coll.UpdateId(id, doc)

	if err != nil {
		return err, nil
	}

	return nil, doc
}

func (m xmodel) Delete(id string) {

}
