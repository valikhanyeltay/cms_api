package db

import (
	"github.com/pkg/errors"
)

func (r *Repository) CreateContentType(contentName, contentDesc string) (id int64, err error) {
	tx := r.db.MustBegin()
	defer tx.Rollback()

	query := `INSERT INTO content_types(content_name, content_desc) VALUES ($1, $2) returning id`

	err = tx.QueryRowx(query, contentName, contentDesc).Scan(&id)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		err = errors.Wrap(err, "tx.Commit()")
	}
	return
}
