package main

import (
	"fmt"
	"net"
)

type Server struct {
	listenAddr string
	listener   net.Listener
	socket     net.Conn
}

func main() {
	server := createServer("localhost:8080")
	server.Start()
}

func createServer(addr string) *Server {
	return &Server{
		listenAddr: addr,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	s.listener = ln
	s.AcceptLoop()
	return nil
}

func (s *Server) AcceptLoop() {
	for {
		socket, err := s.listener.Accept()
		s.socket = socket
		if err != nil {
			fmt.Println("The error is", err.Error())
			continue
		}
		go s.ReadLoop()
	}
}

func (s *Server) ReadLoop() {
	for {
		buffer := make([]byte, 100)
		n, socket_err := s.socket.Read(buffer)
		if socket_err != nil {
			fmt.Println("The socket error is", socket_err.Error())
			return
		}
		msg := string(buffer[:n])
		fmt.Println("Message Received: ", msg)
		if msg == "quit" {
			break
		}
	}
	s.socket.Close()
}
