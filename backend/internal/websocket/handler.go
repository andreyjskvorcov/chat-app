package websocket

import (
	"chat-backend/internal/auth"
	"chat-backend/internal/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var db *sql.DB
var rdb *redis.Client

func Init(database *sql.DB, redisClient *redis.Client) {
	db = database
	rdb = redisClient
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ParseToken(r)
	if err != nil {
		http.Error(w, "Unauthorized", 401)
		return
	}

	room := r.URL.Query().Get("room")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client := &Client{
		conn:   conn,
		userID: userID,
		room:   room,
	}

	addClient(client)

	go listenRedis(room)

	for {
		var msg models.Message

		err := conn.ReadJSON(&msg)
		if err != nil {
			removeClient(client)
			conn.Close()
			break
		}

		msg.UserID = userID
		msg.Room = room

		saveMessage(msg)

		data, _ := json.Marshal(msg)

		rdb.Publish(redis.Ctx, room, data)
	}
}
