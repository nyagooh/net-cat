package tcp

import (
	"bytes"
	"net"
	"sync"
	"testing"
)

// MockConn is a mock implementation of net.Conn for testing purposes.
type MockConn struct {
	net.Conn
	buffer bytes.Buffer
	closed bool // Add this field to track connection closure
}
func (m *MockConn) Write(b []byte) (n int, err error) {
	return m.buffer.Write(b)
}

func (m *MockConn) Read(b []byte) (n int, err error) {
	if m.closed {
		return 0, net.ErrClosed
	}
	return m.buffer.Read(b)
}

// func (m *MockConn) Close() error {
// 	m.closed = true
// 	return nil
// }

func TestServer_broadcast(t *testing.T) {
	type fields struct {
		address      string
		listener     net.Listener
		quitCh       chan struct{}
		clientJoined chan struct{}
		clients      map[net.Conn]string
		history      []string
		mu           sync.Mutex
	}
	type args struct {
		message     string
		excludeConn net.Conn
		logs        bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[net.Conn]string // Expected messages in each connection's buffer
	}{
		{
			name: "Broadcast to all clients except one",
			fields: fields{
				clients: map[net.Conn]string{
					&MockConn{}: "client1",
					&MockConn{}: "client2",
					&MockConn{}: "client3",
				},
			},
			args: args{
				message:     "Hello, World!",
				excludeConn: nil, // No exclusion
				logs:        false,
			},
			want: map[net.Conn]string{},
		},
		{
			name: "Exclude one client from broadcast",
			fields: fields{
				clients: map[net.Conn]string{
					&MockConn{}: "client1",
					&MockConn{}: "client2",
					&MockConn{}: "client3",
				},
			},
			args: args{
				message:     "Hello, World!",
				excludeConn: nil, // Exclude this connection
				logs:        false,
			},
			want: map[net.Conn]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create consistent mock connections
			mockConns := make([]*MockConn, 0, len(tt.fields.clients))
			for conn := range tt.fields.clients {
				mockConns = append(mockConns, conn.(*MockConn))
			}

			// Update the want map with the same mock connections
			for _, mockConn := range mockConns {
				if tt.args.excludeConn != mockConn {
					tt.want[mockConn] = tt.args.message
				} else {
					tt.want[mockConn] = ""
				}
			}

			s := &Server{
				address:      tt.fields.address,
				listener:     tt.fields.listener,
				quitCh:       tt.fields.quitCh,
				clientJoined: tt.fields.clientJoined,
				clients:      tt.fields.clients,
				history:      tt.fields.history,
				mu:           tt.fields.mu,
			}
			s.broadcast(tt.args.message, tt.args.excludeConn, tt.args.logs)

			for conn, expectedMessage := range tt.want {
				mockConn := conn.(*MockConn)
				if got := mockConn.buffer.String(); got != expectedMessage {
					t.Errorf("broadcast() = %v, want %v", got, expectedMessage)
				}
			}
		})
	}
}
