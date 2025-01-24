package database

import (
	"errors"
	"snapp_quera_task/pkg/models"
)

// CreateUser adds a new user
func CreateUser(user *models.User) error {
	return DB.Create(user).Error
}

// GetUserByID fetches a user by ID
func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := DB.First(&user, id).Error
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

// GetAllUsers returns all users
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := DB.Find(&users).Error
	return users, err
}

// UpdateUser updates user details
func UpdateUser(user *models.User) error {
	return DB.Save(user).Error
}

// DeleteUser removes a user
func DeleteUser(id uint) error {
	return DB.Delete(&models.User{}, id).Error
}
