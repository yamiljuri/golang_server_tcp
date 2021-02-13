package protocol

type Protocol interface {
	Match(frame []byte) (bool, error)
	Split(frame []byte) ([][]byte, error)
	Response(frame []byte) ([]byte, error)
	Parser(frame []byte) (interface{}, error)
}
