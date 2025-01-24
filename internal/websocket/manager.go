package websocket

import (
	"log/slog"
	"sync"

	"github.com/gorilla/websocket"
)

// ClientManager manages active WebSocket connections
type ClientManager struct {
	clients map[uint]*websocket.Conn // Maps userID to connection
	mu      sync.Mutex
}

// NewClientManager initializes the WebSocket manager
func NewClientManager() *ClientManager {
	return &ClientManager{
		clients: make(map[uint]*websocket.Conn),
	}
}

// Register adds a new user connection
func (cm *ClientManager) Register(userID uint, conn *websocket.Conn) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.clients[userID] = conn
	slog.Info("User connected", "user_id", userID, "total_users", len(cm.clients))
}

// Unregister removes a user connection
func (cm *ClientManager) Unregister(userID uint, conn *websocket.Conn) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	if _, exists := cm.clients[userID]; exists {
		delete(cm.clients, userID)
		slog.Info("User disconnected", "user_id", userID, "total_users", len(cm.clients))
	}
}

// Broadcast sends a message to all connected users
func (cm *ClientManager) Broadcast(message []byte) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	for userID, conn := range cm.clients {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			slog.Error("Error broadcasting to user", "user_id", userID, "error", err)
			conn.Close()
			delete(cm.clients, userID) // Remove disconnected user
		}
	}
}

// SendPrivateMessage sends a message to a specific user
func (cm *ClientManager) SendPrivateMessage(receiverID uint, message []byte) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	conn, exists := cm.clients[receiverID]
	if !exists {
		slog.Warn("Private message failed, user not connected", "receiver_id", receiverID)
		return
	}
	err := conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		slog.Error("Error sending private message", "receiver_id", receiverID, "error", err)
		conn.Close()
		delete(cm.clients, receiverID) // Remove disconnected user
	}
}
