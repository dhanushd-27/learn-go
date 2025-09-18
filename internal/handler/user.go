package handler

import (
	"context"
	"learn-go/internal/config"
	"learn-go/internal/db/sqlc"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	db     *pgxpool.Pool
	config *config.Config
	query  sqlc.Querier
}

func NewUserHandler(db *pgxpool.Pool, cfg *config.Config, query *sqlc.Queries) *UserHandler {
	return &UserHandler{db: db, config: cfg, query: query}
}

func (h *UserHandler) Me(c echo.Context) error {
	userID := c.Get("user_id")

	user, err := h.query.GetUserById(context.Background(), userID.(int32))

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "User not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"user_id": user.ID,
		"email": user.Email,
		"created_at": user.CreatedAt,
	})
}