package service

import (
	"errors"
	"github.com/valikhanyeltay/cms_api/internal/models"
)

func (s *Service) CreateContentType(contentType *models.ContentType) error {
	if len(contentType.Fields) == 0 {
		return errors.New("no fields in CreateContentType")
	}

	id, err := s.repository.CreateContentType(contentType.ContentName, contentType.ContentDesc)
	if err != nil {
		return err
	}

	if id == 0 {
		return errors.New("invalid content type id returned")
	}

	err = s.repository.AddContentFields(id, contentType.Fields)
	if err != nil {
		return err
	}

	return nil
}
