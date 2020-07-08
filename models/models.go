package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		ID     bson.ObjectId `bson:"_id,omitempty" json:"id"`
		MobNum string        `json:"mobNum"`
	}
)
