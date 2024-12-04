package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

type Message struct {
	from    string
	content []byte
}

type Server struct {
	address  string
	listener net.Listener
	quitCh   chan struct{}
	msgChan  chan Message
	clients  map[net.Conn]string // Map to store client connections and their names
	history  []string            // Slice to store chat history
	mu       sync.Mutex          // Mutex to protect shared resources
}

func NewServer(address string) *Server {
	return &Server{
		address: address,
		quitCh:  make(chan struct{}),
		msgChan: make(chan Message, 100),
		clients: make(map[net.Conn]string),
		history: []string{},
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
	<-s.quitCh //Blocks until a signal is received on quitCh, indicating the server should shut down.
	close(s.msgChan)
	return nil
}

func (s *Server) acceptLoop() {
	for {
		connection, err := s.listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go s.handleNewClient(connection)
	}
}

func (s *Server) handleNewClient(conn net.Conn) {
    defer conn.Close()

    // First send welcome message
    conn.Write([]byte("Welcome to TCP-Chat!\n"))
    
    // Then prompt for name
    conn.Write([]byte("[ENTER YOUR NAME]: "))
    
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

    s.mu.Lock()
    s.clients[conn] = name
    s.mu.Unlock()

    // Send chat history to the new client
    s.sendHistory(conn)

    // Broadcast join message
    joinMsg := fmt.Sprintf("[%s][%s]: %s has joined our chat...\n", 
        time.Now().Format("2006-01-02 15:04:05"), 
        name, 
        name)
    s.broadcast(joinMsg, nil)

    // Start reading messages
    s.readLoop(conn, name)
}

func (s *Server) broadcast(message string, excludeConn net.Conn) {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Add message to history
    s.history = append(s.history, message)

    for conn := range s.clients {
        if conn != excludeConn { // Exclude the specified connection
            if _, err := conn.Write([]byte(message)); err != nil {
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
            s.broadcast(formattedMessage, nil) // Broadcast to all clients including the sender
        }
    }
}
func (s *Server) handleClientDisconnect(conn net.Conn, name string) {
	s.mu.Lock()
	delete(s.clients, conn)
	s.mu.Unlock()
	s.broadcast(fmt.Sprintf("%s has left the chat...\n", name), conn)
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
	server := NewServer(":8989")
	log.Println("Listening on the port :8989")
	log.Fatal(server.StartServer())
}
