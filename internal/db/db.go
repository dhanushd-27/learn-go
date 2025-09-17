package db

import (
	"context"
	"fmt"
	"learn-go/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseService struct {
	cfg *config.Config
	pool *pgxpool.Pool
}

func NewDatabaseService(cfg *config.Config) *DatabaseService {
	return &DatabaseService{cfg: cfg}
}

func (d *DatabaseService) Connect(ctx context.Context) error {
	dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", 
		d.cfg.DB_HOST, d.cfg.DB_PORT, d.cfg.DB_USER, d.cfg.DB_PASSWORD, d.cfg.DB_NAME, d.cfg.SSL_MODE)

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Test the connection
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return fmt.Errorf("failed to ping database: %w", err)
	}

	d.pool = pool
	fmt.Println("Database connection established successfully")
	return nil
}

func (d *DatabaseService) GetPool() *pgxpool.Pool {
	return d.pool
}

func (d *DatabaseService) Close() {
	if d.pool != nil {
		d.pool.Close()
		d.pool = nil
	}
}