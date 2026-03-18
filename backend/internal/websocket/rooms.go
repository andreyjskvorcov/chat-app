package websocket

import "sync"

var rooms = make(map[string]map[*Client]bool)
var mutex = sync.Mutex{}

func addClient(client *Client) {
	mutex.Lock()
	defer mutex.Unlock()

	if rooms[client.room] == nil {
		rooms[client.room] = make(map[*Client]bool)
	}

	rooms[client.room][client] = true
}

func removeClient(client *Client) {
	mutex.Lock()
	defer mutex.Unlock()

	delete(rooms[client.room], client)
}
