package connection

type Connection interface {
	GetId() int64
	Read()
	Write(buffer []byte)
	Close()
}
