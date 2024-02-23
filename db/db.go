package db

import (
	"database/sql"
	"github.com/EraldCaka/github-authentication-app/util"
	_ "github.com/lib/pq"
)

type database struct {
	db *sql.DB
}

func NewPostgres() (*database, error) {
	db, err := sql.Open("postgres", util.DB_URL)
	if err != nil {
		return nil, err
	}

	return &database{db: db}, nil
}

func (d *database) GetDB() *sql.DB {
	return d.db
}

func (d *database) Ping() error {
	return d.db.Ping()
}

func (d *database) Close() {
	d.db.Close()
}
