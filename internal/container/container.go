package container

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	db *pgxpool.Pool
}

// func NewContainer() (*Container, error) {
// 	// Write code here to connect with db
// }

func (dc *Container) GetDB() *pgxpool.Pool {
	return dc.db
}