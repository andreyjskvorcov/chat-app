package redis

import (
	"chat-backend/internal/models"
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return rdb
}

func saveMessage(msg models.Message) {
	_, err := db.Exec(
		"INSERT INTO messages (user_id, room, text) VALUES ($1, $2, $3)",
		msg.UserID, msg.Room, msg.Text,
	)
	if err != nil {
		log.Println(err)
	}
}

func listenRedis(room string) {
	sub := rdb.Subscribe(redis.Ctx, room)

	for msg := range sub.Channel() {
		broadcast(room, []byte(msg.Payload))
	}
}
