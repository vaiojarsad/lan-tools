package database

import (
	"database/sql"
	"errors"
)

func Create() error {
	db, err := Open()
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	var errs []error
	for _, t := range ddls {
		_, err = db.Exec(t)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}
