package database

import (
	"log"
	"chatroom/pkg/models"
)

// RunMigrations applies database migrations
func RunMigrations() {
	err := DB.AutoMigrate(&models.User{}, &models.Message{})
	if err != nil {
		log.Fatal("❌ Failed to run migrations:", err)
	}
	log.Println("✅ Migrations applied successfully")
}
