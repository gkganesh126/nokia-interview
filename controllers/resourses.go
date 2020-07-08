package controllers

import (
	_ "net/http/pprof"

	"github.com/gkganesh126/nokia-interview/models"
)

type (
	// For Get
	UserResources struct {
		Data []models.User `json:"data"`
	}
	// For Post/Put
	UserResource struct {
		Data models.User `json:"data"`
	}
)
