package tcp

import (
	"net"
	"strings"
	"sync"
	"testing"
)

func TestServer_handleClientDisconnect(t *testing.T) {
	mockConn := &MockConn{}
	server := &Server{
		clients: make(map[net.Conn]string),
		mu:      sync.Mutex{},
	}

	// Add mockConn to the server's clients
	server.clients[mockConn] = "TestUser"

	// Create another mock connection to capture broadcast messages
	anotherMockConn := &MockConn{}
	server.clients[anotherMockConn] = "AnotherUser"

	// Call handleClientDisconnect
	server.handleClientDisconnect(mockConn, "TestUser")

	// Verify that mockConn is removed from the server's clients
	if _, exists := server.clients[mockConn]; exists {
		t.Errorf("Expected client to be removed from server.clients")
	}

	// Verify that the leave message was broadcasted to anotherMockConn
	expectedMessage := "TestUser has left the chat...\n"
	if !strings.Contains(anotherMockConn.buffer.String(), expectedMessage) {
		t.Errorf("Expected message %q, got %q", expectedMessage, anotherMockConn.buffer.String())
	}
}
