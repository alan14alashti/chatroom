package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"io"
)

func Login() {
	var email, password string
	fmt.Print("Enter Email: ")
	fmt.Scan(&email)
	fmt.Print("Enter Password: ")
	fmt.Scan(&password)

	credentials := map[string]string{"email": email, "password": password}
	data, _ := json.Marshal(credentials)

	resp, err := http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("❌ Error logging in:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result map[string]string
	json.Unmarshal(body, &result)

	token, exists := result["token"]
	if !exists {
		fmt.Println("❌ Login failed!")
		return
	}

	// Save token to file
	_ = os.WriteFile("token.txt", []byte(token), 0600)
	fmt.Println("✅ Login successful! Token saved.")
}
