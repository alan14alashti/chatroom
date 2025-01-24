package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"chatroom/config"
	"chatroom/internal/api"
	"chatroom/internal/database"
	"chatroom/internal/middleware"
	"chatroom/internal/websocket"
	"chatroom/pkg/logger"
	"syscall"
	"time"
)

func main() {
	// Load environment variables
	cfg := config.LoadConfig()
	log := logger.InitLogger()

	// Connect to the database
	database.ConnectDatabase(cfg)

	// Run database migrations
	database.RunMigrations()

	// Initialize WebSocket manager
	wsManager := websocket.NewClientManager()

	// HTTP server with WebSocket and API routes
	mux := http.NewServeMux()

	// WebSocket route with JWT authentication
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.HandleConnections(wsManager, w, r)
	})

	// Auth & User management routes
	mux.HandleFunc("/register", api.RegisterUserHandler)
	mux.HandleFunc("/login", api.LoginUserHandler)

	// Protected routes
	mux.HandleFunc("/user", middleware.JWTMiddleware(api.GetUserHandler))
	mux.HandleFunc("/users", middleware.JWTMiddleware(api.GetAllUsersHandler))
	mux.HandleFunc("/delete-user", middleware.JWTMiddleware(api.DeleteUserHandler))

	server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: mux,
	}

	// Run the server in a separate goroutine
	go func() {
		log.Info("🚀 Server started", "port", cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("❌ Server error", "error", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Info("⚠️ Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Error("❌ Shutdown error", "error", err)
	}
	log.Info("✅ Server stopped gracefully")
}
