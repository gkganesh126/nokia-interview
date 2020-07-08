package data

import (
	"github.com/gkganesh126/nokia-interview/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct {
	C *mgo.Collection
}

func (r *UserRepository) Create(user *models.User) error {
	obj_id := bson.NewObjectId()
	user.ID = obj_id
	err := r.C.Insert(&user)
	return err
}

func (r *UserRepository) GetAll() []models.User {
	var users []models.User
	iter := r.C.Find(nil).Iter()
	result := models.User{}
	for iter.Next(&result) {
		users = append(users, result)
	}
	return users
}

func (r *UserRepository) Delete(id bson.ObjectId) error {
	err := r.C.Remove(bson.M{"_id": id})
	return err
}

func (r *UserRepository) Update(user *models.User) error {
	return r.C.Update(bson.M{"_id": user.ID}, bson.M{"$set": bson.M{"mobnum": user.MobNum}})
}
