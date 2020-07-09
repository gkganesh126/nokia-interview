package controllers

import (
	"github.com/gkganesh126/nokia-interview/controllers/cache"
)

type Storage interface {
	Get(key string) []byte
	Set(key string, content []byte)
	GetAll() []cache.Item
}

var StorageTemp cache.Storage
