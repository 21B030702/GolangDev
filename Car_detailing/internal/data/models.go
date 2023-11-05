package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	CarDetails CarDetailModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		CarDetails: CarDetailModel{DB: db},
	}
}
