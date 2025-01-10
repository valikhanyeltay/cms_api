package db

import (
	"github.com/valikhanyeltay/cms_api/internal/models"
	"github.com/pkg/errors"
)

func (r *Repository) AddContentFields(contentId int64, fields []models.Field) error {
	tx := r.db.MustBegin()
	defer tx.Rollback()

	query := `INSERT INTO content_fields (content_id, field_name, field_type, default_value) 
			  VALUES ($1, $2, $3, $4)`

	for _, field := range fields {
		_, err := tx.Exec(query, contentId, field.Name, field.Type, field.DefaultValue)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err := tx.Commit()
	if err != nil {
		return errors.Wrap(err, "tx.Commit()")
	}
	
	return nil
}
