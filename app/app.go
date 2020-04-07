package app

import (
	"github.com/yamiljuri/server_tcp/config"
	"github.com/yamiljuri/server_tcp/database"
	"github.com/yamiljuri/server_tcp/server"
)

func init() {
	config.LoadEnv()
}

func Start() {
	//Inicializamos MongoDb
	database.Default()
	server := server.NewServer()
	server.Run()
}
