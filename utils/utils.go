package utils

import (
	"reflect"

	"github.com/yamiljuri/server_tcp/connection"
)

func Remove(slice []connection.Connection, position int64) []connection.Connection {

	result := make([]connection.Connection, len(slice)-1)
	for _, connection := range slice {
		if connection.GetId() != position {
			result = append(result, connection)
		}
	}
	return result
}

func ParserObjectByTag(tag string, object interface{}) map[string]map[string]reflect.Value {
	value := reflect.ValueOf(object)
	nameObject := value.Type().Name()
	result := make(map[string]map[string]reflect.Value)
	resultField := make(map[string]reflect.Value)
	for i := 0; i < value.Type().NumField(); i++ {
		if valueField := value.Type().Field(i).Tag.Get(tag); valueField != "" && valueField != "-" {
			resultField[valueField] = value.Field(i)
		}
	}
	result[nameObject] = resultField
	return result
}
