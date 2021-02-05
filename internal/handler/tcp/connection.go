package tcp

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/yamiljuri/server_tcp/internal/core/protocol"
)

type Connection interface {
	GetId() int64
	Read()
	Write(buffer []byte)
	Close()
}

type connection struct {
	IdConnection int64
	OutChannel   chan<- Message
	InChannel    chan []byte
	IsConnect    bool
	protocol     protocol.Protocol
	conn         net.Conn
}

func NewConnection(id int64, conn net.Conn, out chan<- Message) Connection {
	client := connection{conn: conn, OutChannel: out}
	client.InChannel = make(chan []byte)
	client.IsConnect = true
	client.IdConnection = id
	client.Read()
	return &client
}

func (c *connection) GetId() int64 {
	return c.IdConnection
}

func (c *connection) Read() {
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

				if frames, err := c.protocol.Split(buffer); err == nil {
					for _, frame := range frames {
						report, err := c.protocol.Parser(frame)
						if err != nil {
							fmt.Printf("Error Parser: %s", err)
						}
					}
				}

				message := Message{}
				message.idClientAplication = c.IdConnection
				message.idTypeMessage = _Message_Type_Send
				message.message = []byte(buffer)
				c.OutChannel <- message
				log.Printf("Length Buffer %d", n)
			}
		}
	}()
}

func (c *connection) Write(buffer []byte) {
	msg := fmt.Sprintf("%s => Messages broadcast: %s", time.Now().Format("15:04:05"), buffer)
	_, err := c.conn.Write([]byte(msg))
	if err != nil {
		log.Print(err)
	}
}

func (c *connection) Close() {
	log.Printf("Se cierra el cliente %d", c.IdConnection)
	message := Message{}
	message.idClientAplication = c.IdConnection
	message.idTypeMessage = _Message_Type_Close
	c.IsConnect = false
	close(c.InChannel)
	c.OutChannel <- message
	c.conn.Close()
}
