package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yamiljuri/server_tcp/cmd/app/dependencies"
)

func Start(dependencies dependencies.Dependencies) {

	//gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	Routes(engine, dependencies)

	log.Fatal(engine.Run(":8080"))
}
