package database

import (
	"errors"
	"chatroom/pkg/models"
)

// SaveMessage stores a message in the database
func SaveMessage(senderID, receiverID uint, content string) error {
	message := models.Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
	}
	return DB.Create(&message).Error
}

// GetMessages retrieves all messages (public + private)
func GetMessages(userID uint) ([]models.Message, error) {
	var messages []models.Message
	err := DB.Where("receiver_id = ? OR receiver_id = 0", userID).Find(&messages).Error
	return messages, err
}

// GetPrivateMessages retrieves private messages between two users
func GetPrivateMessages(userID1, userID2 uint) ([]models.Message, error) {
	var messages []models.Message
	err := DB.Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
		userID1, userID2, userID2, userID1).Find(&messages).Error

	if err != nil {
		return nil, errors.New("failed to retrieve messages")
	}

	return messages, nil
}
