package websocket

import (
	"log/slog"
	"sync"

	"github.com/gorilla/websocket"
)

// ClientManager manages WebSocket clients
type ClientManager struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

// NewClientManager initializes the WebSocket manager
func NewClientManager() *ClientManager {
	return &ClientManager{
		clients: make(map[*websocket.Conn]bool),
	}
}

// Register adds a new client
func (cm *ClientManager) Register(client *websocket.Conn) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.clients[client] = true
	slog.Info("Client connected", "total", len(cm.clients))
}

// Unregister removes a client
func (cm *ClientManager) Unregister(client *websocket.Conn) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	delete(cm.clients, client)
	slog.Info("Client disconnected", "total", len(cm.clients))
}
