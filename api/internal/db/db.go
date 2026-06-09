package db

import (
	"flowmoney/api/internal/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDb(cfg config.DB) (*sqlx.DB, error) {
	connStr := fmt.Sprintf(" host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Pass, cfg.Name, cfg.SslMode)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
