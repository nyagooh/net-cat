package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

const maxClients = 10

type Server struct {
	address      string
	listener     net.Listener
	quitCh       chan struct{}
	clientJoined chan struct{} 
	clients      map[net.Conn]string
	history      []string
	mu           sync.Mutex
}

func NewServer(address string) *Server {
	return &Server{
		address:      address,
		quitCh:       make(chan struct{}),
		clientJoined: make(chan struct{}),
		clients:      make(map[net.Conn]string),
		history:      []string{},
	}
}

func (s *Server) StartServer() error {
	listen, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	s.listener = listen
	defer listen.Close()
	go s.acceptLoop()
	<-s.quitCh
	return nil
}

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
	file,err := os.ReadFile("file.txt")
	if err!= nil {
        log.Println("Error opening file:", err)
    }
	conn.Write([]byte(file))

	nameBuf := make([]byte, 1024)
	n, err := conn.Read(nameBuf)
	if err != nil || n == 0 {
		return
	}
	name := strings.TrimSpace(string(nameBuf[:n]))
	if name == "" {
		conn.Write([]byte("Name cannot be empty. Disconnecting...\n"))
		conn.Close()

	}

	s.mu.Lock()
	s.clients[conn] = name
	for _, ch := range s.clients {
		if ch == name {
			log.Println("Please input another name:name already in use")
		}
	}
	clientCount := len(s.clients) 
	s.mu.Unlock()

	
	s.sendHistory(conn)

	if clientCount > 1 {
		joinMsg := fmt.Sprintf("%s has joined our chat...\n",
			name,
		)
		s.broadcast(joinMsg, conn, true) 
	}


	s.readLoop(conn, name)
}

func (s *Server) broadcast(message string, excludeConn net.Conn, logs bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !logs {
		s.history = append(s.history, message)
	}
	for conn := range s.clients {
		if conn != excludeConn { 
			_, err := conn.Write([]byte(message))
			if err != nil {
				log.Println("Error writing to connection:", err)
			}
		}
	}
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

func (s *Server) sendHistory(conn net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, msg := range s.history {
		if _, err := conn.Write([]byte(msg)); err != nil {
			log.Println("Error writing history to connection:", err)
		}
	}
}

func main() {
	port := "8989"
	if len(os.Args) > 2 {
		return
	} else if len(os.Args) == 2 {
		port = os.Args[1]
	}
	server := NewServer(":" + port)
	log.Println("Listening on the port", port)
	log.Fatal(server.StartServer())
}
