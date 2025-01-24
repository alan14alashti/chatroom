package websocket

import (
	"log/slog"
	"net/http"
	"chatroom/internal/auth"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// HandleConnections handles WebSocket connections
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

		slog.Info("Message received", "user_id", userID, "message", string(msg))
	}
}
