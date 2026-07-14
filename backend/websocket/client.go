package websocket

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(c *gin.Context) {

	conn, err := Upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		println("Upgrade failed:", err.Error())
		return
	}
		println("Client connected")

	AlertHub.Mutex.Lock()

	AlertHub.Clients[conn] = true

	AlertHub.Mutex.Unlock()

	defer func() {
		println("Client disconnected")
		AlertHub.Mutex.Lock()

		delete(AlertHub.Clients, conn)

		AlertHub.Mutex.Unlock()

		conn.Close()

	}()

	for {

		if _, _, err := conn.ReadMessage(); err != nil {

			break
		}
	}
}