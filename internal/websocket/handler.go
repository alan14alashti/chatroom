package websocket

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"chatroom/internal/auth"
	"chatroom/internal/database"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type MessagePayload struct {
	ReceiverID uint   `json:"receiver_id"` // 0 for public messages
	Content    string `json:"content"`
}

// HandleConnections manages WebSocket connections with JWT authentication
func HandleConnections(manager *ClientManager, w http.ResponseWriter, r *http.Request) {
	tokenString := r.URL.Query().Get("token")
	userID, err := auth.ValidateJWT(tokenString)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("Error upgrading connection", "error", err)
		return
	}
	defer conn.Close()

	manager.Register(userID, conn)
	defer manager.Unregister(userID, conn)

	slog.Info("User connected", "user_id", userID)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			slog.Error("Error reading message", "error", err)
			break
		}

		// Parse the received message
		var payload MessagePayload
		if err := json.Unmarshal(msg, &payload); err != nil {
			slog.Error("Invalid message format", "error", err)
			continue
		}

		// Store message in the database
		if err := database.SaveMessage(userID, payload.ReceiverID, payload.Content); err != nil {
			slog.Error("Failed to save message", "error", err)
			continue
		}

		// Send the message to the appropriate user(s)
		if payload.ReceiverID == 0 {
			// Broadcast to all users
			manager.Broadcast(msg)
		} else {
			// Send private message
			manager.SendPrivateMessage(payload.ReceiverID, msg)
		}
	}
}
