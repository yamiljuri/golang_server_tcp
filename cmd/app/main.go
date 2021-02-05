package main

import (
	"github.com/yamiljuri/server_tcp/cmd/app/dependencies"
	tcpHandler "github.com/yamiljuri/server_tcp/cmd/app/tcp"
)

func main() {
	dep := dependencies.New().Initialize()
	//api.Start(dep)
	tcpHandler.Start(dep)
}
