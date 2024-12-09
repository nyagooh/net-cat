package main

import (
	"fmt"
	"log"
	"os"
	"netcat/tcp"
)

func main() {
	port := "8989"
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	} else if len(os.Args) == 2 {
		port = os.Args[1]
	}
	server := tcp.NewServer(":" + port)
	if server == nil {
		log.Fatalf("[USAGE]: ./TCPChat $port")
		return
	}
	fmt.Println("Listening on the port:", port)
	log.Fatal(server.StartServer())
}
