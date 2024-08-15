package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func formatHTTPResponse(response string) string {
	// Format the response as an HTTP response
	// This is done so that the client can understand the response and display it correctly
	return fmt.Sprintf("HTTP/1.1 200 OK\nContent-Type: text/plain\nContent-Length: %d\n\n%s", len(response), response)
}

func processRequest(request string) string {
	// Process the incoming request - here we parse the request & log path, method, headers, body, etc.
	// The request is in the format: "METHOD /path HTTP/1.1\nHeader1: Value1\nHeader2: Value2\n\nBody"
	requestLines := strings.Split(request, "\n")
	// The first line contains the method, path, and HTTP version
	firstLine := strings.Split(requestLines[0], " ")
	method := firstLine[0]
	path := firstLine[1]
	httpVersion := firstLine[2]
	fmt.Printf("Method: %s\nPath: %s\nHTTP Version: %s\n", method, path, httpVersion)
	return formatHTTPResponse("Hello, World!")
}

func handleConnection(conn net.Conn) {
	// Defer closing the connection until the function returns
	defer conn.Close()

	// Create a buffer to store the incoming request
	// Read only the first 1024 bytes of the request
	buffer := make([]byte, 1024)

	// Read the incoming request into the buffer - this is a blocking call
	n, err := conn.Read(buffer)
	if err != nil {
		log.Panicf("Failed to read data: %s", err)
	}
	fmt.Printf("# Bytes Read: %d\n", n)
	// Parse the incoming request
	request := string(buffer[:n])

	// Process the request
	response := processRequest(request)

	// Write the response back to the client - this is a blocking call
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Panicf("Failed to write data: %s", err)
	}

}

func main() {
	// Start a TCP server on port 8090
	listener, err := net.Listen("tcp", ":8090")

	if err != nil {
		log.Panicf("Failed to start server: %s", err)
	}

	// Defer closing the listener until the function returns
	defer listener.Close()

	// Run the server indefinitely (or until an error occurs) - we can stop the server by sending a SIGINT signal (Ctrl+C)
	// Without this loop, the server would only accept a single connection and then exit
	for {

		// Accept incoming connections - this is a blocking call (it will wait until a connection is made)
		conn, err := listener.Accept()

		if err != nil {
			log.Panicf("Failed to accept connection: %s", err)
		}

		// Handle the incoming connection
		// 🔥 Multi-Threaded Server: We process the request in a separate goroutine, this allows the server to handle multiple connections parallely
		go handleConnection(conn)

		fmt.Print("\n\n")
	}
}
