package cli

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Welcome to Chat CLI!")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Start Chat")
	fmt.Println("4. View Online Users")
	fmt.Println("5. View Chat History")
	fmt.Println("6. Exit")

	var choice int
	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		Register()
	case 2:
		Login()
	case 3:
		StartChat()
	case 4:
		GetOnlineUsers()
	case 5:
		GetChatHistory()
	case 6:
		fmt.Println("Exiting...")
		return
	default:
		log.Println("Invalid choice, try again.")
	}
}
