package container

import (
	"context"
	"learn-go/internal/config"
	"learn-go/internal/db"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	db     *pgxpool.Pool
	Config *config.Config
}

func NewContainer() (*Container, error) {
	// Load configuration
	config, err := config.Load()
	if err != nil {
		return nil, err
	}

	// Create database service and connect
	dbService := db.NewDatabaseService(config)
	if err := dbService.Connect(context.Background()); err != nil {
		return nil, err
	}

	return &Container{
		db:     dbService.GetPool(),
		Config: config,
	}, nil
}

func (dc *Container) GetDB() *pgxpool.Pool {
	return dc.db
}
