package websocket

import "encoding/json"

func Broadcast(data interface{}) {

	bytes, _ := json.Marshal(data)

	AlertHub.Mutex.Lock()
	defer AlertHub.Mutex.Unlock()

	for client := range AlertHub.Clients {
		client.WriteMessage(1, bytes)
	}
}