package server

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/yamiljuri/server_tcp/config"
	"github.com/yamiljuri/server_tcp/connection"
	"github.com/yamiljuri/server_tcp/database"
	"github.com/yamiljuri/server_tcp/models/frame"
	messageType "github.com/yamiljuri/server_tcp/models/message"
	"github.com/yamiljuri/server_tcp/utils"
)

type Server struct {
	Address               string
	Type                  string
	receiveMessageChannel chan messageType.MessageServer
	Connections           []connection.Connection
}

func NewServer() Server {
	server := Server{Type: "tcp", Address: fmt.Sprintf("%s:%s", config.Getenv("SERVER_HOST"), config.Getenv("SERVER_PORT"))}
	return server
}

func (s *Server) receiveMessage() {
	for {
		select {
		case message := <-s.receiveMessageChannel:
			switch message.GetTypeMessage() {
			case messageType.Send:
				log.Printf("Receive %d => %s", message.GetIdClientAplication(), message.GetMessage())
				frame := frame.Frame{
					DeviceId: int(message.GetIdClientAplication()),
					Frame:    bytes.Trim(message.GetMessage(), string(byte(0))),
					Date:     time.Now(),
					Type:     frame.Input,
				}
				database.DB.Insert(frame)
			case messageType.Close:
				log.Printf("Close %d", message.GetIdClientAplication())
				s.Connections = utils.Remove(s.Connections, message.GetIdClientAplication())

			}
		}
	}
}

func (s *Server) readStdIn() {
	go func() {
		for {
			input, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				log.Printf("Error al leer std, %v", err)
			}
			connections := s.Connections
			for _, v := range connections {
				if v != nil {
					v.Write([]byte(input))
				}
			}

		}
	}()
}
func (s *Server) Run() {
	server, err := net.Listen(s.Type, s.Address)
	if err != nil {
		log.Fatalf("Error Start Server TCP , %v", err)
	}
	defer server.Close()
	s.receiveMessageChannel = make(chan messageType.MessageServer)
	s.readStdIn()
	go s.receiveMessage()
	i := 0
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Printf("Error Connetion Accept , %v", conn)
		}

		connection := connection.NewClientConnection(int64(i), conn, s.receiveMessageChannel)
		if connection == nil {
			log.Printf("Error al crear una connection %v", connection)
		}
		s.Connections = append(s.Connections, connection)
		i++
	}
}
