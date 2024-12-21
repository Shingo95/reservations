package dbrepo

import (
	"database/sql"

	"github.com/elorenzotti/bookings/internal/config"
	"github.com/elorenzotti/bookings/internal/repository"
)

type postgresdbrepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {

	return &postgresdbrepo{
		App: a,
		DB:  conn,
	}
}
