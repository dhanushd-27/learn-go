package handler

import (
	"context"
	"net/http"

	"learn-go/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	db     *pgxpool.Pool
	config *config.Config
}

func NewAuthHandler(db *pgxpool.Pool, cfg *config.Config) *AuthHandler {
	return &AuthHandler{db: db, config: cfg}
}

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Signup(c echo.Context) error {
	var req SignupRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	// Placeholder: demonstrate dependency usage
	if err := h.db.Ping(context.Background()); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database unavailable"})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"status":  "ok",
		"message": "signup successful (placeholder)",
		"email":   req.Email,
		"port":    h.config.PORT,
	})
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	// Placeholder: normally verify password, generate token, etc.
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "ok",
		"message": "login successful (placeholder)",
		"email":   req.Email,
	})
}
