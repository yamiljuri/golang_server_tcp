package example

import "github.com/yamiljuri/server_tcp/internal/core/protocol"

type Example struct {
}

func New() protocol.Protocol {
	return &Example{}
}

func (e *Example) Match(frame []byte) (bool, error) {
	return false, nil
}

func (e *Example) Split(frame []byte) ([][]byte, error) {
	return [][]byte{frame}, nil
}

func (e *Example) Response(frame []byte) ([]byte, error) {
	return nil, nil
}

func (e *Example) Parser(frame []byte) (interface{}, error) {
	return nil, nil
}
