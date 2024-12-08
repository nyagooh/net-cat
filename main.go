package main

import (
	"fmt"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	// Read data from the connection
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	// Print the data received
	fmt.Println("Received message:", string(buffer[:n]))
}

func main() {

	//error handling
	if len(os.Args) != 1 {
		log.Fatal("\nUsage: go run .\n", )
	}

	// Print welcome message
	fmt.Println("Server is starting...")

	// Determine the port from environment variable, default to 8080 if not set
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	// Start listening for incoming connections on port 8080 as the default port
	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatalf("\nError starting server: %v", err)
	}
	defer listener.Close()
	log.Infof("\nServer listening on localhost:%s", port)

	// Infinite loop to accept and handle incoming client connections
	for {
		// Accept a new incoming connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}
