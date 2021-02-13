package dependencies

import (
	"os"
)

type Dependencies interface {
	Initialize() Dependencies
}

type dependencies struct {
}

func New() Dependencies {
	return &dependencies{}
}

func (d *dependencies) Initialize() Dependencies {
	switch os.Getenv("ENVIROMENT") {
	case "test":
		d.initEnvTesting()
	default:
		d.initEnvProduction()
	}
	return d
}
