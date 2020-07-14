package cache

import (
	"testing"
)

func TestStorage(t *testing.T) {
	storage := NewStorage()

	storage.Get("hey")

	tables := []struct {
		key      string
		value    string
		expected string
	}{
		{"hey", "dude", "dude"},
		{"wow", "wow", "wow"},
	}

	for _, table := range tables {
		storage := NewStorage()
		storage.Set(table.key, []byte(table.value))
		got := storage.Get(table.key)
		if string(got) != table.expected {
			t.Errorf("expected: %s, got: %s", table.expected, string(got))
		}
	}
}
