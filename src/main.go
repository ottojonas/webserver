package main

import (
	"fmt"
	"log"
	"net/http"

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

func main() {
    http.HandleFunc("/", helloHandler)
    http.HandleFunc("/ws", wsHandler)
	
	port := ":3001"
	fmt.Printf("Server is running on http://localhost%s\n", port)
    log.Fatal(http.ListenAndServe(port, nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "wagwan")
}
