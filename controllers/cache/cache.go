package cache

import (
	"sync"
)

// Item is a cached reference
type Item struct {
	Content []byte
}

//Storage mecanism for caching strings in memory
type Storage struct {
	items map[string]Item
	mu    *sync.RWMutex
}

//NewStorage creates a new in memory storage
func NewStorage() *Storage {
	return &Storage{
		items: make(map[string]Item),
		mu:    &sync.RWMutex{},
	}
}

//Get a cached content by key
func (s Storage) Get(key string) []byte {
	s.mu.RLock()
	defer s.mu.RUnlock()

	item := s.items[key]
	return item.Content
}

//Set a cached content by key
func (s Storage) Set(key string, content []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.items[key] = Item{
		Content: content,
	}
}

func (s Storage) GetAll() []Item {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var items []Item
	for _, data := range s.items {
		items = append(items, data)
	}
	return items
}
