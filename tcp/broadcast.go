package tcp

import (
	"log"
	"net"
)

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

func (s *Server) sendHistory(conn net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, msg := range s.history {
		if _, err := conn.Write([]byte(msg)); err != nil {
			log.Println("Error writing history to connection:", err)
		}
	}
}
