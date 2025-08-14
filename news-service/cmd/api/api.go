package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"news-service/cmd/middleware"
	"news-service/internal/domain/article/delivery"
	"news-service/internal/domain/article/usecase"
	"news-service/package/config"
	"news-service/package/logger"

	deliveryTags "news-service/internal/domain/tags/delivery"
	ucTags "news-service/internal/domain/tags/usecase"

	deliveryArticleVersion "news-service/internal/domain/article-version/delivery"
	ucArticleVersion "news-service/internal/domain/article-version/usecase"

	"news-service/metrics"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/urfave/cli/v2"

	_ "news-service/docs"

	"news-service/package/validator"

	echoMiddlerware "github.com/labstack/echo/v4/middleware"
)

const CmdServeHTTP = "serve-http"

type HTTP struct{
	usecase usecase.IArticle
	ucTags ucTags.ITags
	ucArticleVersion ucArticleVersion.IArticleVersionUsecase
	conf *config.Config
	v *validator.Validator
}

func (h HTTP) ServeAPI(c *cli.Context) error  {
	if err := logger.SetLogger(); err != nil {
		log.Printf("error logger %v", err)
	}
	// Register metrics
	metrics.Register()
	e := echo.New();
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/health-check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok!")
	})

	e.Use(echoMiddlerware.CORSWithConfig(echoMiddlerware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))
	
	// Configurable rate limiter
	ipLimiter := middleware.RateLimiterMiddleware(
		cast.ToInt(h.conf.RateLimitMaxRequest), 
		time.Duration(cast.ToInt(h.conf.RateLimitInterval)) * time.Second, 
		cast.ToFloat64(h.conf.RateLimitJitter),
	) 
	e.Use(ipLimiter.Middleware())

	articleAPI := e.Group("/api/v1")
	articleAPI.Use(middleware.LoggerMiddleware)
	articleAPI.Use(middleware.MonitoringMiddleware)

	delivery.NewArticleHTTP(articleAPI, h.usecase, h.v)
	deliveryTags.NewTagHTTP(articleAPI, h.ucTags, h.v)
	deliveryArticleVersion.NewArticleVersionHTTP(articleAPI, h.ucArticleVersion, h.v)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := e.Start(fmt.Sprintf(":%v", h.conf.AppPort)); err != nil {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	return nil	
}

func ServeAPI(conf *config.Config, v *validator.Validator, 
	usecase usecase.IArticle,
	avUsecase ucArticleVersion.IArticleVersionUsecase,
	) []*cli.Command {
	h := &HTTP{conf: conf, usecase: usecase, ucArticleVersion: avUsecase, v: v}
	return []*cli.Command{
		{
			Name: CmdServeHTTP,
			Usage: "Serve News Service",
			Action: h.ServeAPI,
		},
	}
}