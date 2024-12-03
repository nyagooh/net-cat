package main

import (
	"fmt"
	"net"
)

func Client() {
	conn, err := net.Dial("tcp", "localhost:8080")
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
