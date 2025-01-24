package database

import (
	"log"
	"chatroom/pkg/models"
)

func RunMigrations() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("❌ Failed to run migrations:", err)
	}
	log.Println("✅ Migrations applied successfully")
}
