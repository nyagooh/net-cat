package main

import (
	"fmt"
	"net"
	"os"
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
		fmt.Println("Usage: go run .")
		os.Exit(1)
	}

	// Print welcome message
	fmt.Println("Server is starting...")

	// Start listening for incoming connections on port 8080
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

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
