package tcp

const (
	_Message_Type_Send      = 1
	_Message_Type_Close     = 2
	_Message_Type_Heardbeat = 3
)

type Message struct {
	id                 int64
	idClientAplication int64
	idTypeMessage      int
	message            []byte
	report             interface{}
}
