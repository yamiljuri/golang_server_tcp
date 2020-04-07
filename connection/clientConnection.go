package connection

import (
	"fmt"
	"log"
	"net"
	"time"

	messageType "github.com/yamiljuri/server_tcp/models/message"
)

type ClientConnection struct {
	IdConnection int64
	OutChannel   chan<- messageType.MessageServer
	InChannel    chan []byte
	IsConnect    bool
	conn         net.Conn
}

func NewClientConnection(id int64, conn net.Conn, out chan<- messageType.MessageServer) Connection {
	client := ClientConnection{conn: conn, OutChannel: out}
	client.InChannel = make(chan []byte)
	client.IsConnect = true
	client.IdConnection = id
	client.Read()
	return &client
}

func (c *ClientConnection) GetId() int64 {
	return c.IdConnection
}

func (c *ClientConnection) Read() {
	go func() {
		defer c.Close()
		buffer := make([]byte, 1024)
		for {
			if c.IsConnect {
				n, err := c.conn.Read(buffer)
				if err != nil {
					log.Printf("Error al leer: %v", err)
					break
				}
				message := messageType.MessageServer{}
				message.SetIdClientAplication(c.IdConnection)
				message.SetTypeMessage(messageType.Send)
				message.SetMessage([]byte(buffer))
				c.OutChannel <- message
				log.Printf("Length Buffer %d\n\r", n)
			}
		}
	}()
}

func (c *ClientConnection) Write(buffer []byte) {
	msg := fmt.Sprintf("%s => Messages broadcast: %s\n\r", time.Now().Format("15:04:05"), buffer)
	_, err := c.conn.Write([]byte(msg))
	if err != nil {
		log.Print(err)
	}
}

func (c *ClientConnection) Close() {
	log.Printf("Se cierra el cliente %d \n\r", c.IdConnection)
	message := messageType.MessageServer{}
	message.SetIdClientAplication(c.IdConnection)
	message.SetTypeMessage(messageType.Close)
	c.IsConnect = false
	close(c.InChannel)
	c.OutChannel <- message
	c.conn.Close()
}
