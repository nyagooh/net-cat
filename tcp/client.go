package netcat

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func (s *Server) acceptLoop() {
	for {
		connection, err := s.listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			connection.Close()
		}

		s.mu.Lock()
		if len(s.clients) >= maxClients {
			s.mu.Unlock()
			connection.Write([]byte("Server is full. Please try again later.\n"))
			connection.Close()
		}
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
	conn.Write([]byte("[ENTER YOUR NAME]:"))

	nameBuf := make([]byte, 1024)
	n, err := conn.Read(nameBuf)
	if err != nil || n == 0 {
		return
	}
	name := strings.TrimSpace(string(nameBuf[:n]))
	if name == "" {
		conn.Write([]byte("Name cannot be empty. Disconnecting...\n"))
		conn.Close()
		return
	}
	for _, v := range name {
		if !isLetter(v) {
            conn.Write([]byte("Name can only contain letters. Disconnecting...\n"))
            conn.Close()
            return
        }
	}

	s.mu.Lock()
	var clientCount int

	s.clients[conn] = name
	clientCount = len(s.clients)
	s.mu.Unlock()
	s.sendHistory(conn)

	if clientCount > 1 {
		joinMsg := fmt.Sprintf("%s has joined our chat...\n", name)
		s.broadcast(joinMsg, conn, true)
	}

	s.readLoop(conn, name)
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
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
		if message != "" {
			timestamp := time.Now().Format("2006-01-02 15:04:05")
			formattedMessage := fmt.Sprintf("[%s][%s]: %s\n", timestamp, name, message)
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
