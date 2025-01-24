package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	SenderID   uint   `json:"sender_id"`
	ReceiverID uint   `json:"receiver_id"` // 0 if it's a public message
	Content    string `json:"content"`
}
