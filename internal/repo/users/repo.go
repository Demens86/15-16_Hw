package users

import "github.com/jmoiron/sqlx"

type Repo struct {
	db *sqlx.DB
	// log
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}
