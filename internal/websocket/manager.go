package websocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type ClientManager struct {
	clients map[uint]*websocket.Conn
	mu      sync.Mutex
}

func (cm *ClientManager) Register(userID uint, conn *websocket.Conn) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.clients[userID] = conn
}

func (cm *ClientManager) Unregister(userID uint, conn *websocket.Conn) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	delete(cm.clients, userID)
}

