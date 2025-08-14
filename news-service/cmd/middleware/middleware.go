package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"news-service/package/config"
	"news-service/package/utils"

	"github.com/labstack/echo/v4"
)

type contextKey string
const (
    ctxUserID contextKey = "user_id"
    ctxRoleID contextKey = "role"
)

var cfg = config.NewConfig()
// Middleware untuk Writer (role = 3)
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        tokenStr := c.Request().Header.Get("Authorization")
        if tokenStr == "" {
            return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
        }
        fmt.Println("debug", cfg.JwtSecretKey)

        tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)
        userID, role, err := utils.ValidateJWT(tokenStr, cfg)
        if err != nil {
            return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
        }

        if role > 3 {
            return echo.NewHTTPError(http.StatusForbidden, "Forbidden: Writer access required")
        }

        c.Set("user_id", userID)
        c.Set("role", role)
        return next(c)
    }
}


// Middleware untuk Editor (role = 1 atau 2)
func AuthEditorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        tokenStr := c.Request().Header.Get("Authorization")
        if tokenStr == "" {
            return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
        }

        tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)
        userID, role, err := utils.ValidateJWT(tokenStr, cfg) // role ikut divalidasi
        if err != nil {
            return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
        }

        if role != 1 && role != 2 {
            return echo.NewHTTPError(http.StatusForbidden, "Forbidden: Editor access required")
        }

        c.Set("user_id", userID)
        c.Set("role", role)
        return next(c)
    }
}