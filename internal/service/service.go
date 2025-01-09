package service

import (
	"fmt"
	"github.com/valikhanyeltay/cms_api/internal/db"
	"github.com/valikhanyeltay/cms_api/internal/models"
	"reflect"
	"strings"
)

type Service struct {
	structRegistry map[string]reflect.Type
	repository     *db.Repository
}

func NewService(repo *db.Repository) *Service {
	service := &Service{
		repository:     repo,
		structRegistry: loadStructRegistry(nil),
	}

	return service
}

func (s *Service) RegisterContent(contentType string, payload map[string]interface{}) {
	fields := make([]reflect.StructField, 0, len(payload))
	for key, val := range payload {
		fields = append(fields, reflect.StructField{
			Name: capitalFirstLetter(key),
			Type: reflect.TypeOf(val),
			Tag:  reflect.StructTag(fmt.Sprintf(`json:"%s"`, key)),
		})
	}

	contentStruct := reflect.StructOf(fields)

	s.structRegistry[contentType] = contentStruct
}

func capitalFirstLetter(s string) string {
	return strings.ToTitle(s[:1]) + s[1:]
}

func loadStructRegistry(contentTypes []models.ContentType) map[string]reflect.Type {
	structRegistry := make(map[string]reflect.Type, len(contentTypes))

	for _, contentType := range contentTypes {
		structFields := make([]reflect.StructField, 0, len(contentType.Fields))

		for _, field := range contentType.Fields {
			structFields = append(structFields, reflect.StructField{
				Name: capitalFirstLetter(field.Name),
				Type: determineType(field.Type),
				Tag:  reflect.StructTag(fmt.Sprintf(`json:"%s"`, field.Name)),
			})
		}

		structRegistry[contentType.ContentName] = reflect.StructOf(structFields)
	}

	return structRegistry
}

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
