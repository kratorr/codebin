package postgresql

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

func NewPostgresClient(ctx context.Context, host, port, user, password, dbname string) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s/%s?user=%s&password=%s", host, port, dbname, user, password)
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
