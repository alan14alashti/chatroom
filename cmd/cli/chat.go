package cli

import (
	"fmt"
	"os"
	"log"
	"net/url"
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
