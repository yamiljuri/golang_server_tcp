package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/yamiljuri/server_tcp/internal/core/report"
)

type Server interface {
	Run()
}

type server struct {
	Address                string
	Port                   int
	receivedMessageChannel chan Message
	Connections            []Connection
	Service                report.Report
}

func NewServer(address string, port int) Server {
	return &server{Address: address, Port: port}
}

func (s *server) Run() {
	addressPortListener := fmt.Sprintf("%s:%d", s.Address, s.Port)
	socketListen, err := net.Listen("tcp", addressPortListener)
	if err != nil {
		log.Fatalf("Error Start Server TCP , %v", err)
	}
	defer socketListen.Close()
	s.receivedMessageChannel = make(chan Message)
	s.readStdIn()
	go s.receiveMessage()
	i := 0
	for {
		conn, err := socketListen.Accept()
		if err != nil {
			log.Printf("Error Connetion Accept , %v", conn)
		}

		connection := NewConnection(int64(i), conn, s.receivedMessageChannel)
		if connection == nil {
			log.Printf("Error al crear una connection %v", connection)
		}
		s.Connections = append(s.Connections, connection)
		i++
	}
}

func (s *server) receiveMessage() {
	for {
		select {
		case message := <-s.receivedMessageChannel:
			switch message.idTypeMessage {
			case _Message_Type_Send:
				log.Printf("Receive %d => %s", message.idClientAplication, message.message)
				s.Service.Save(message)
			case _Message_Type_Close:
				log.Printf("_Message_Type_Close %d", message.idClientAplication)
				s.Connections = s.remove(s.Connections, message.idClientAplication)
			}
		}
	}
}

func (s *server) readStdIn() {
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

func (s *server) remove(slice []Connection, position int64) []Connection {
	var result []Connection
	for _, connection := range slice {
		if connection != nil && connection.GetId() != position {
			result = append(result, connection)
		}
	}
	return result
}
