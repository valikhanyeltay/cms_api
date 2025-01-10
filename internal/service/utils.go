package service

import (
	"reflect"
	"strings"
)

func determineType(typeStr string) reflect.Type {
	switch typeStr {
	case "string":
		return reflect.TypeOf("")
	case "int":
		return reflect.TypeOf(0)
	case "float64":
		return reflect.TypeOf(float64(0))
	case "bool":
		return reflect.TypeOf(false)
	case "[]interface{}":
		return reflect.TypeOf([]interface{}{})
	case "map[string]interface{}":
		return reflect.TypeOf(make(map[string]interface{}))
	case "nil":
		return reflect.TypeOf(nil)
	default:
		return reflect.TypeOf(nil)
	}
}

func capitalFirstLetter(s string) string {
	return strings.ToTitle(s[:1]) + s[1:]
}
