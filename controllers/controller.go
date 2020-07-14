package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/gkganesh126/nokia-interview/controllers/cache"
	db "github.com/gkganesh126/nokia-interview/db-ops"
	"github.com/segmentio/kafka-go"
)

type Storage interface {
	Get(key string) []byte
	Set(key string, content []byte)
	GetAll() []cache.Item
}

var StorageCache cache.Storage

func ReloadCacheFromDb() {
	// read from db
	mongoContext := NewContext()
	defer mongoContext.Close()
	c := mongoContext.NokiaDbCollection("users")
	repo := &db.UserRepository{c}
	// Get all users form repository
	users := repo.GetAll()

	// to produce messages
	topic := "test"
	partition := 0

	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	for _, user := range users {
		conn.WriteMessages(
			kafka.Message{Key: []byte(user.ID), Value: []byte(user.MobNum)},
		)
	}

	conn.Close()

	time.Sleep(2000 * time.Millisecond)

	conn, _ = kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	//	b := make([]byte, 10e3) // 10KB max per message
	for {
		msg, err := batch.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(string(msg.Key), " ", string(msg.Value))
		StorageCache.Set(string(msg.Key), []byte(msg.Key))
	}

	batch.Close()
	conn.Close()

	items := StorageCache.GetAll()

	for key, item := range items {
		fmt.Println("key:", key, "\t value:", string(item.Content))
	}

}
