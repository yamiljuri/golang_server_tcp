package config

import (
	"bytes"
	"io/ioutil"
	"log"
)

var (
	config = map[string]map[string]string{
		"server": {
			"SERVER_HOST": "0.0.0.0",
			"SERVER_PORT": "5555",
		},
		"database": {
			"MONGO_DB_DRIVER":   "mongodb",
			"MONGO_DB_HOST":     "127.0.0.1",
			"MONGO_DB_PORT":     "27017",
			"MONGO_DB_USER":     "",
			"MONGO_DB_PASSWORD": "",
			"MONGO_DB_DATABASE": "",
		},
	}
)

func LoadEnv() {

	variablesLoadFile := loadFile(".env")
	if config != nil {
		for configSection, parameters := range config {
			for parameter, _ := range parameters {
				if valueEnv := variablesLoadFile[parameter]; valueEnv != "" {
					config[configSection][parameter] = valueEnv
				}
			}
		}
	}
}

func loadFile(file string) map[string]string {
	varEnv := make(map[string]string)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Error Enviroment Load File %v", err)
	}
	dataSplit := bytes.Split(data, []byte{0xA})
	for _, value := range dataSplit {
		if !bytes.HasPrefix(value, []byte{0x23}) && bytes.Contains(value, []byte{0x3d}) {
			parameter := bytes.Split(value, []byte{0x3d})
			varEnv[string(parameter[0])] = string(parameter[1])
		}
	}
	return varEnv
}

func Getenv(key string) string {
	valueParameter := ""
	if config != nil {
		for _, parameters := range config {
			if valueEnv := parameters[key]; valueEnv != "" {
				valueParameter = valueEnv
			}
		}
	}
	return valueParameter
}
