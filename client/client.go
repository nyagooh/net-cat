package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func Client() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("dial error", err)
		return
	}
	defer conn.Close()

	fmt.Println("Enter message (type 'exit' to quit):")
	message := ""
	for {

		//allow the user to exit from the
		fmt.Scan(&message)
		if message == "exit" || message == "quit" || strings.HasPrefix(message, "bye") {
			break
		}
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}
}

func main() {

	//error checking
	if len(os.Args) != 1 {
		log.Fatal("\nUsage: go run .\n")
		os.Exit(1)
	}
	Client()
}
