package middleware

import (
	"auth-service/package/config"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)


type client struct {
	Requests int
	Expiry   time.Time
}

var (
	mu         sync.Mutex
	clients    = make(map[string]*client)
	timeWindow = time.Minute  // durasi window
)
 
func RatelimitMidleware(cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			mu.Lock()
			defer mu.Unlock()

			ip := c.RealIP()
			now := time.Now()
			timeExpired := now.Add(time.Millisecond * time.Duration(cast.ToInt(cfg.RateLimitThreshold)))

			cl, exists := clients[ip]
			if !exists || now.After(cl.Expiry) {
				clients[ip] = &client{Requests: 1, Expiry: now.Add(timeWindow)}
			} else {
				if cl.Requests >= cast.ToInt(cfg.RatelimitMaxRetry) {
					retryAfter := cl.Expiry.Sub(timeExpired).Milliseconds()
					c.Response().Header().Set("Retry-After", fmt.Sprintf("%.0f", retryAfter))
					return c.JSON(http.StatusTooManyRequests, map[string]string{
						"error": "Rate limit exceeded. Try again later.",
					})
				}
				cl.Requests++
			}
				return next(c)
		}
	}
}