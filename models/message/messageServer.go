package message

const (
	Send      = 1
	Close     = 2
	Heardbeat = 3
)

type MessageServer struct {
	id                 int64
	idClientAplication int64
	idTypeMessage      int
	message            []byte
}

func (m *MessageServer) SetId(id int64) {
	m.id = id
}
func (m *MessageServer) SetIdClientAplication(idClientAplication int64) {
	m.idClientAplication = idClientAplication
}
func (m *MessageServer) SetTypeMessage(typeMessage int) {
	m.idTypeMessage = typeMessage
}
func (m *MessageServer) SetMessage(message []byte) {
	m.message = message
}

func (m *MessageServer) GetId() int64 {
	return m.id
}
func (m *MessageServer) GetIdClientAplication() int64 {
	return m.idClientAplication
}
func (m *MessageServer) GetTypeMessage() int {
	return m.idTypeMessage
}
func (m *MessageServer) GetMessage() []byte {
	return m.message
}
