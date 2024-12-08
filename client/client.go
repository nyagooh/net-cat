package main

import (
	"fmt"
	"net"
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
	Client()
}
