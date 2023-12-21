package databases

import (
	"log"
	"lolyshop/config"

	"github.com/jmoiron/sqlx"
)

func DbConnect(cfg config.IDbConfig) *sqlx.DB {
	db, err := sqlx.Connect("pgx", cfg.Url())
	if err != nil {
		log.Fatalf("connect to database failed: %v", err)
	}

	db.DB.SetMaxOpenConns(cfg.MaxOpenConns())

	return db
}
