package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type Client struct {
	conn     net.Conn
	name     string
	serverCh chan string
}

func NewClient(address string) (*Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	client := &Client{
		conn:     conn,
		serverCh: make(chan string),
	}

	return client, nil
}

func (c *Client) handleServerMessages() {
    reader := bufio.NewReader(c.conn)
    for {
        message, err := reader.ReadString('\n')
        if err != nil {
            log.Println("Error reading from server:", err)
            c.serverCh <- "error"
            return
        }
        // Directly print the message received from the server
        fmt.Print(message)
    }
}

func (c *Client) sendMessage() {
    reader := bufio.NewReader(os.Stdin)
    for {
        text, err := reader.ReadString('\n')
        if err != nil {
            log.Println("Error reading input:", err)
            continue
        }
        text = strings.TrimSpace(text)
        if text == "" {
            continue
        }

        // Send the raw message text to the server
        _, err = fmt.Fprintf(c.conn, "%s\n", text)
        if err != nil {
            log.Println("Error sending message:", err)
            return
        }
    }
}

func (c *Client) setName() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	name = strings.TrimSpace(name)
	if name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	c.name = name
	_, err = fmt.Fprintf(c.conn, "%s\n", name)
	return err
}

func main() {
	address := "localhost:8989" // Change this to the server's address if needed
	client, err := NewClient(address)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer client.conn.Close()

	if err := client.setName(); err != nil {
		log.Fatal("Error setting name:", err)
	}

	go client.handleServerMessages()
	client.sendMessage()
}
