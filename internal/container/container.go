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

func NewContainer(ctx context.Context, cfg *config.Config) (*Container, error) {
	dbService := db.NewDatabaseService(cfg)
	if err := dbService.Connect(ctx); err != nil {
		return nil, err
	}

	return &Container{
		db:     dbService.GetPool(),
		Config: cfg,
	}, nil
}

func (dc *Container) GetDB() *pgxpool.Pool {
	return dc.db
}

func (dc *Container) Close() {
	if dc.db != nil {
		dc.db.Close()
	}
}
