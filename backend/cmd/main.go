package main

import (
	"log"
	"net/http"

	"chat-backend/internal/database"
	"chat-backend/internal/redis"
	"chat-backend/internal/websocket"
)

func main() {
	db := database.InitPostgres()
	rdb := redis.InitRedis()

	websocket.Init(db, rdb)

	http.HandleFunc("/ws", websocket.HandleConnections)

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
