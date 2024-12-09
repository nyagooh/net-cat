package netcat

import (
	"net"
	"sync"
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
