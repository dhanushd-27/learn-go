package db

import "github.com/jackc/pgx/v5/pgxpool"
type DBConnect interface {
	Connect() (*pgxpool.Pool, error)
}

type DatabaseService {
	config string
}
