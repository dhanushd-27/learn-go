package middleware

import (
	"learn-go/internal/config"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTCookieMiddleware(cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("token")
			if err != nil || cookie.Value == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing JWT token")
			}

			token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token signing method")
				}
				return []byte(cfg.JWT_SECRET), nil
			})

			if err != nil || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid JWT token")
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid JWT claims")
			}

			c.Set("user", token)

			if idFloat, ok := claims["user_id"].(float64); ok {
				c.Set("user_id", int32(idFloat))
			} else {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid user_id in token")
			}

			if email, ok := claims["email"].(string); ok {
				c.Set("email", email)
			}

			return next(c)
		}
	}
}
