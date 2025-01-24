package api

import (
	"encoding/json"
	"net/http"
	"chatroom/internal/auth"
	"chatroom/internal/database"
	"chatroom/pkg/models"
	"strconv"
)

// RegisterUserHandler - Creates a new user with a hashed password
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Hash the password before saving
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	// Save the user
	if err := database.CreateUser(&user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// LoginUserHandler - Authenticates user and returns JWT token
func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Find user by email
	user, err := database.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Verify password
	if !auth.CheckPassword(user.Password, req.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := auth.GenerateJWT(user.ID)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// GetUserHandler - Retrieves user details (protected route)
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := database.GetUserByID(uint(id))
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Hide password in response
	user.Password = ""

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// GetAllUsersHandler - Retrieves all users (protected route)
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := database.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	// Hide passwords in response
	for i := range users {
		users[i].Password = ""
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// DeleteUserHandler - Deletes a user (protected route)
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := database.DeleteUser(uint(id)); err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
