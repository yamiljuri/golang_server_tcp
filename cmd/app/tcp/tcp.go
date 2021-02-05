package tcp

import (
	"github.com/yamiljuri/server_tcp/cmd/app/dependencies"
	"github.com/yamiljuri/server_tcp/internal/handler/tcp"
)

func Start(dependencies dependencies.Dependencies) {
	tcp.NewServer("0.0.0.0", 9999).Run()
}
