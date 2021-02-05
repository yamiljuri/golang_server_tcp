package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yamiljuri/server_tcp/cmd/app/dependencies"
)

func Routes(engine *gin.Engine, dependencies dependencies.Dependencies) {
	//engine.Use(dependencies.CageHttp.Interceptor)
	//engine.POST("/cage", dependencies.CageHttp.Post)
}
