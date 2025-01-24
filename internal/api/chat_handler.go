package api

import (
	"encoding/json"
	"net/http"
	"chatroom/internal/auth"
	"chatroom/internal/database"
)

// GetChatHistoryHandler retrieves chat history for a logged-in user
func GetChatHistoryHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	userID, err := auth.ValidateJWT(tokenString)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	messages, err := database.GetChatHistory(userID)
	if err != nil {
		http.Error(w, "Failed to retrieve chat history", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
