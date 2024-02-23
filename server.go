package server

import (
	"fmt"
	"net"
)

type server struct {
	address  string
	listener net.Listener
}

type handlerType func(conn net.Conn, server *server) error

func CreateServer(host string, port uint16, networkType string) (*server, error) {
	address := fmt.Sprintf("%s:%d", host, port)
	listener, err := net.Listen(networkType, address)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s server listening on %s \n", networkType, address)
	serverInstance := server{address, listener}
	return &serverInstance, nil
}

func (s *server) AcceptConnections(handler handlerType) {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			conn.Close()
			fmt.Println("Could not accept connection", err)
			continue
		}
		go handler(conn, s)
	}
}
