package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/golang/glog"
	"github.com/gkganesh126/nokia-interview/common"
	db "github.com/gkganesh126/nokia-interview/db-ops"
	"gopkg.in/mgo.v2"
)

// Handler for HTTP Get - "/users"
// Returns all User documents
func GetUsers(w http.ResponseWriter, r *http.Request) {
	glog.Info("At GetUsers...\n")
	// Create new context
	context := NewContext()
	defer context.Close()
	c := context.NokiaDbCollection("users")
	repo := &db.UserRepository{c}
	// Get all users form repository
	users := repo.GetAll()
	j, err := json.Marshal(UserResources{Data: users})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Post - "/users"
// Create a new Showtime document
func CreateUser(w http.ResponseWriter, r *http.Request) {
	glog.Info("At CreateUser...\n")
	var dbResource UserResource
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dbResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid User db", 500)
		return
	}
	user := &dbResource.Data
	glog.Info("user: ", user)
	context := NewContext()
	defer context.Close()
	c := context.NokiaDbCollection("users")
	// Create User
	repo := &db.UserRepository{c}
	repo.Create(user)
	// Create response db
	j, err := json.Marshal(dbResource)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Delete - "/users/{id}"
// Delete a User document by id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	glog.Info("At DeleteUser\n")

	// Create new context
	context := NewContext()
	defer context.Close()

	var dbResource UserResource
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dbResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid User db", 500)
		return
	}
	id := &dbResource.Data.ID
	// Create new context

	c := context.NokiaDbCollection("users")

	// Remove user by id
	repo := &db.UserRepository{c}
	err = repo.Delete(*id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
			return
		}
	}

	// Send response back
	w.WriteHeader(http.StatusNoContent)
}

func UsersUpdate(w http.ResponseWriter, r *http.Request) {
	glog.Info("At UpdateUser...\n")

	var dbResource UserResource
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dbResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid User db", 500)
		return
	}
	user := &dbResource.Data
	// Create new context
	context := NewContext()
	defer context.Close()
	c := context.NokiaDbCollection("users")
	// Create User
	repo := &db.UserRepository{c}
	repo.Update(user)
	// Create response db
	j, err := json.Marshal(dbResource)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
