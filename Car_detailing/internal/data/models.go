package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	CarDetails CarDetailModel
	Tokens     TokenModel
	Users      UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		CarDetails: CarDetailModel{DB: db},
		Tokens:     TokenModel{DB: db},
		Users:      UserModel{DB: db},
	}
}
