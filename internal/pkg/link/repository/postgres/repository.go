package postgres

import (
	link "github.com/White-AK111/shortener/internal/pkg"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) link.Repository {
	ret := repository{
		db: db,
	}

	return ret
}
