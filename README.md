# 🚀 Chatroom Project (CLI & WebSocket)

A **real-time chat application** built with **Go, WebSockets, and PostgreSQL**, featuring:  
✅ **User Authentication** (Register & Login with JWT)  
✅ **Public & Private Messaging** (via WebSockets)  
✅ **User Presence Tracking** (Who's online?)  
✅ **Chat History Retrieval** (Stored in PostgreSQL)  
✅ **Dockerized Deployment** (Runs with `docker-compose`)  

## 🛠 Setup & Installation

### **1️⃣ Clone the Repository**
```bash
git clone https://github.com/your-repo/chatroom.git
cd chatroom
2️⃣ Configure Environment Variables
Create a .env file and add:

POSTGRES_DB=chatroom
POSTGRES_USER=alan
POSTGRES_PASSWORD=secretpassword
POSTGRES_PORT=5422
POSTGRES_HOST=postgres
SERVER_PORT=8080
3️⃣ Start the Application with Docker

docker-compose up --build
Starts PostgreSQL
Builds & runs the Go application
4️⃣ Use the CLI for Chat

go run cmd/cli/main.go

💡 Features
👤 User Authentication
Register a new user (/register)
Login to get a JWT token (/login)
Secure endpoints using JWT
💬 WebSocket Real-Time Messaging
Public chat (messages visible to all users)
Private chat (direct messages between users)
Stores messages in PostgreSQL
👥 User Presence Tracking
Tracks who is online (/online-users)
📜 Chat History Retrieval
Fetch previous messages (/chat-history)
📦 Fully Dockerized
PostgreSQL & Go services run in Docker containers
Uses docker-compose for easy deployment

🚀 Usage Examples
🔹 Register a User

curl -X POST http://localhost:8080/register \
     -d '{"name":"alan","email":"alan14alashti@gamil.com","password":"123456"}' \
     -H "Content-Type: application/json"
🔹 Login & Get JWT Token

curl -X POST http://localhost:8080/login \
     -d '{"email":"alice@example.com","password":"123"}' \
     -H "Content-Type: application/json"
🔹 Start Chat via CLI

go run cmd/cli/main.go
Send public messages: Hello everyone!
Send private messages: @2 Hello, User 2!
Check online users: Choose Option 4
View chat history: Choose Option 5
📌 Next Steps
🔹 Add Message Read Receipts
🔹 Implement User Typing Indicators
🔹 Deploy using Docker + Kubernetes + NATS server implementation
