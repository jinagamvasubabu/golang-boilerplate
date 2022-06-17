package model

import (
	"gopkg.in/mgo.v2/bson"
)

type MongoBookEntity struct {
	Id     bson.ObjectId `bson:"_id,omitempty"`
	Title  string        `bson:"title"`
	Author string        `bson:"author"`
	Desc   string        `bson:"desc"`
}
