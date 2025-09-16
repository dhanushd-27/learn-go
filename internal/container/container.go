package container

import (
	"learn-go/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	db     *pgxpool.Pool
	Config *config.Config
}

func NewContainer() (*Container, error) {
	// Write code here to connect with db
	config, err := config.Load()
	if err != nil {
		return nil, err
	}

	return &Container{
		db:     nil,
		Config: config,
	}, nil
}

func (dc *Container) GetDB() *pgxpool.Pool {
	return dc.db
}
