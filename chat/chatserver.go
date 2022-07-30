package chat

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

var publicChat chan string = make(chan string)

type server struct {
	connections []*net.Conn
}

func (s *server) Broadcast() {
	for message := range publicChat {
		for _, conn := range s.connections {
			fmt.Fprintln(*conn, message)
		}
	}
}

func (s *server) HandleNewConnection(conn net.Conn) {
	defer conn.Close()
	finished := make(chan struct{})

	fmt.Fprintln(conn, "Welcome to the server!")
	clientName := conn.RemoteAddr().String()

	publicChat <- fmt.Sprintf("Client %s has joined the chat room\n", clientName)

	go func(c chan<- struct{}) {
		inputMessage := bufio.NewScanner(conn)
		for inputMessage.Scan() {
			publicChat <- fmt.Sprintf("%s: %s", clientName, inputMessage.Text())
		}

		c <- struct{}{}
	}(finished)

	<-finished
	publicChat <- fmt.Sprintf("Client %s left the chat room\n", clientName)
}

func RunServer() {
	flag.Parse()
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	serv := server{connections: make([]*net.Conn, 0)}

	fmt.Println("Server listening on localhost:8080")

	go serv.Broadcast()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("There was an error:", err)
		}
		serv.connections = append(serv.connections, &conn)
		go serv.HandleNewConnection(conn)
	}
}
