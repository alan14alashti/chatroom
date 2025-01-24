package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Register() {
	var name, email, password string
	fmt.Print("Enter Name: ")
	fmt.Scan(&name)
	fmt.Print("Enter Email: ")
	fmt.Scan(&email)
	fmt.Print("Enter Password: ")
	fmt.Scan(&password)

	userData := map[string]string{
		"name":     name,
		"email":    email,
		"password": password,
	}

	data, _ := json.Marshal(userData)
	resp, err := http.Post("http://localhost:8080/register", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("❌ Error registering:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		fmt.Println("✅ Registration successful!")
	} else {
		fmt.Println("❌ Registration failed!")
	}
}
