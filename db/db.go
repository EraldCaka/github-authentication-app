package db

import (
	"context"
	"github.com/EraldCaka/github-authentication-app/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"sync"
)

type Postgres struct {
	db *pgxpool.Pool
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

func NewPGInstance(ctx context.Context) (*Postgres, error) {
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, util.DB_URL)
		if err != nil {
			log.Println("Unable to connect to Postgres Db: %w", err)
			return
		}
		pgInstance = &Postgres{db}
	})

	return pgInstance, nil
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.db.Close()
}
