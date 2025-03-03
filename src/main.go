package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"webserver/src/api"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader {
	CheckOrigin: func(r * http.Request) bool { return true }, 
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err) 
		return 
	}
	defer conn.Close()
	for {
		messageType, message, err := conn.ReadMessage() 
		if err != nil {
			log.Println("Read error: ", err) 
			break
		}
		log.Printf("Received: %s", message)
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Println("Write error:", err) 
			break 
		}
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API Endpoint") 
}

func main() {
    http.HandleFunc("/", helloHandler)
    http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/api/users", getAllUsersHandler)
	http.HandleFunc("/api/users/{id}", getUserHandler)
	http.HandleFunc("/api/createUser", createUserHandler) 
	
	port := ":3001"
	fmt.Printf("Server is running on http://localhost%s\n", port)
    log.Fatal(http.ListenAndServe(port, nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "wagwan")
}

func getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(api.Users)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return 
	}
	idStr := parts[3]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return 
	}
	for _, user := range api.Users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return 
		}
	}
	http.NotFound(w, r)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser api.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	newUser.ID = len(api.Users) + 1
	api.Users = append(api.Users, newUser)
	w.Header().Set("Content-Type", "application/json") 
	json.NewEncoder(w).Encode(newUser)
}
