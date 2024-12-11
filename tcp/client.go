package tcp

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func (s *Server) AcceptLoop() {
	for {
		connection, err := s.listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue // Skip to the next iteration if there's an error
		}

		s.mu.Lock()
		if len(s.clients) > maxClients {
			s.mu.Unlock()
			connection.Write([]byte("Server is full. Please try again later.\n"))
			connection.Close()
			continue // Skip to the next iteration if the server is full
		}
		s.clients[connection] = "" // Add the connection to the clients map
		s.mu.Unlock()

		go s.handleNewClient(connection)
	}
}

func (s *Server) handleNewClient(conn net.Conn) {
	defer conn.Close()
	file, err := os.ReadFile("file.txt")
	if err != nil {
		log.Println("Error opening file:", err)
	}
	conn.Write([]byte(file))

	for {
		conn.Write([]byte("[ENTER YOUR NAME]:"))

		nameBuf := make([]byte, 1024)
		n, err := conn.Read(nameBuf)
		if err != nil || n == 0 {
			return
		}
		name := strings.TrimSpace(string(nameBuf[:n]))
		if name == "" {
			conn.Write([]byte("Name cannot be empty. Disconnecting...\n"))
			return
		}

		// Check for special characters
		for _, v := range name {
			if !isLetterOrDigit(v) {
				conn.Write([]byte("Name can only contain letters and numbers. Disconnecting...\n"))
				return
			}
		}

		s.mu.Lock()
		nameTaken := false
		for _, existingName := range s.clients {
			if existingName == name {
				nameTaken = true
				break
			}
		}
		s.mu.Unlock()

		if nameTaken {
			conn.Write([]byte("Name is already taken. Please choose a different name.\n"))
		} else {
			s.mu.Lock()
			s.clients[conn] = name
			clientCount := len(s.clients)
			s.mu.Unlock()

			s.sendHistory(conn)

			if clientCount > 1 {
				joinMsg := fmt.Sprintf("%s has joined our chat...\n", name)
				s.broadcast(joinMsg, conn, true)
			}

			s.readLoop(conn, name)
			break
		}
	}
}

func isLetterOrDigit(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func (s *Server) readLoop(conn net.Conn, name string) {
	buffer := make([]byte, 512)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("read error:", err)
			s.handleClientDisconnect(conn, name)
			return
		}
		message := strings.TrimSpace(string(buffer[:n]))
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		formattedMessage := fmt.Sprintf("[%s][%s]: %s\n", timestamp, name, message)

		if message == "" {
			// Send the empty message only to the sender
			conn.Write([]byte("\033[F\033[K"))
			conn.Write([]byte(formattedMessage))
		} else {
			// Broadcast non-empty messages to all clients
			conn.Write([]byte("\033[F\033[K"))
			s.broadcast(formattedMessage, nil, false)
		}
	}
}

func (s *Server) handleClientDisconnect(conn net.Conn, name string) {
	s.mu.Lock()
	delete(s.clients, conn)
	s.mu.Unlock()
	s.broadcast(fmt.Sprintf("%s has left the chat...\n", name), conn, true)
}
