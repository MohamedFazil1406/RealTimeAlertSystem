package websocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	Clients map[*websocket.Conn]bool

	Mutex sync.Mutex
}

var AlertHub = &Hub{
	Clients: make(map[*websocket.Conn]bool),
}