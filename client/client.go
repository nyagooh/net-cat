package main

import (
	"fmt"
	"net"
	"os"
)

func Client() {

	//set the port
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	conn, err := net.Dial("tcp", "localhost:"+port)
	fmt.Println("enter message")
	message := ""
	for {

		fmt.Scan(&message)
		if err != nil {
			fmt.Println("dial error", err)
		}
		conn.Write([]byte(message))
	}
}

func main() {
	Client()
}
