package service

import (
	"github.com/valikhanyeltay/cms_api/internal/db"
	"log"
	"reflect"
)

type Service struct {
	structRegistry map[string]reflect.Type
	repository     *db.Repository
}

func NewService(repo *db.Repository) *Service {
	// upload data from DB on start
	structTypes, err := repo.GetContentTypes()
	if err != nil {
		log.Panic(err)
	}
	service := &Service{
		repository:     repo,
		structRegistry: loadStructRegistry(structTypes),
	}

	return service
}
