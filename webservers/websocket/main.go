package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrade the HTTP connection to a WebSocket connection
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins (you can add more security here)
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// http.ResponseWriter is used to write a response back to the client
	// http.Request object contains information about the incoming request.

	// 1. Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade to WebSocket: %s", err)
		return
	}
	defer conn.Close()

	// 2. Persistent connection is established; Now, we can read and write messages.
	for {
		// 3. ReadMessage reads the next message from the client
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Failed to read message: %s", err)
			break
		}

		// Log the received message
		fmt.Printf("Received: %s\n", message)

		// 4. WriteMessage sends a message back to the client
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Printf("Failed to write message: %s", err)
			break
		}
	}
}

func main() {
	// Register the handleWebSocket function to handle WebSocket connections on the /ws endpoint
	http.HandleFunc("/ws", handleWebSocket)

	// Start the server on port 8090
	log.Println("WebSocket server started on :8090")

	// ListenAndServe starts an HTTP server and blocks the current goroutine (this is the same as .Accept() in the TCP server)
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
