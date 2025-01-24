package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"log"
	"net/url"
	"net/http"
	"github.com/gorilla/websocket"
)

// StartChat connects to the WebSocket server and sends/receives messages
func StartChat() {
	// Load token
	tokenBytes, err := os.ReadFile("token.txt")
	if err != nil {
		fmt.Println("❌ Please log in first.")
		return
	}
	token := string(tokenBytes)

	// Connect to WebSocket server
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws", RawQuery: "token=" + token}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("❌ Failed to connect to WebSocket:", err)
	}
	defer conn.Close()

	fmt.Println("✅ Connected to chat! Type messages and press Enter to send.")
	go readMessages(conn)

	for {
		var message string
		fmt.Print("> ")
		fmt.Scanln(&message)

		if message == "/exit" {
			fmt.Println("Exiting chat...")
			return
		}

		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			fmt.Println("❌ Error sending message:", err)
			break
		}
	}
}

// readMessages listens for incoming messages
func readMessages(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("❌ Disconnected from server.")
			break
		}
		fmt.Println("\n📩 New message:", string(message))
		fmt.Print("> ")
	}
}

// GetOnlineUsers fetches the list of online users
func GetOnlineUsers() {
	resp, err := http.Get("http://localhost:8080/online-users")
	if err != nil {
		fmt.Println("❌ Error fetching online users:", err)
		return
	}
	defer resp.Body.Close()

	var users []uint
	json.NewDecoder(resp.Body).Decode(&users)
	fmt.Println("👥 Online Users:", users)
}

// GetChatHistory fetches the user's chat history
func GetChatHistory() {
	tokenBytes, err := os.ReadFile("token.txt")
	if err != nil {
		fmt.Println("❌ Please log in first.")
		return
	}
	token := string(tokenBytes)

	req, _ := http.NewRequest("GET", "http://localhost:8080/chat-history", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("❌ Error fetching chat history:", err)
		return
	}
	defer resp.Body.Close()

	var messages []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&messages)

	fmt.Println("📜 Chat History:")
	for _, msg := range messages {
		fmt.Printf("📩 [%v] %v -> %v: %v\n", msg["created_at"], msg["sender_id"], msg["receiver_id"], msg["content"])
	}
}