package models

type Message struct {
	UserID string `json:"user_id"`
	Room   string `json:"room"`
	Text   string `json:"text"`
}
