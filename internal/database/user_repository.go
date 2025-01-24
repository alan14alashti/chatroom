package database

import (
	"errors"
	"chatroom/pkg/models"

	"gorm.io/gorm"
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail fetches a user by email (needed for login)
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// GetAllUsers returns all users (excluding soft-deleted)
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := DB.Find(&users).Error
	return users, err
}

// UpdateUser updates only provided fields
func UpdateUser(id uint, updates map[string]interface{}) error {
	err := DB.Model(&models.User{}).Where("id = ?", id).Updates(updates).Error
	return err
}

// DeleteUser performs a soft delete
func DeleteUser(id uint) error {
	err := DB.Delete(&models.User{}, id).Error
	return err
}
