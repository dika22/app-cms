package delivery

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	avRepository "news-service/internal/domain/article-version/repository"
	"news-service/internal/domain/article/repository"
	cacheRepo "news-service/internal/domain/article/repository/cache"
	"news-service/internal/domain/article/usecase"
	authorRepository "news-service/internal/domain/author/repository"
	tagsRepository "news-service/internal/domain/tags/repository"

	"news-service/package/config"
	"news-service/package/connection/cache"

	"github.com/stretchr/testify/assert"

	// database "news-service/package/connection/database/sqlite"
	database "news-service/package/connection/database"
	"news-service/package/validator"

	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func SetupTestHTTP(t *testing.T)  {
	e = echo.New()
	dbConf := config.NewDatabase()
  	conf := config.NewConfig()
  	dbConn := database.NewDatabase("wdb", dbConf)
  	cacheConf := config.NewCache()
  	cache := cache.NewRedis("cache-cms", cacheConf)

  	articleRepo := repository.NewsRepository(dbConn, cache)
  	authorRepo := authorRepository.NewAuthorRepository(dbConn)
	tagsRepo := tagsRepository.NewTagsRepository(dbConn)
	avRepo := avRepository.NewArticleVersionRepository(dbConn)
  	cacheArticleRepo := cacheRepo.NewCacheRepository(cache)
  	validate := validator.NewValidator()
  	usecase := usecase.NewsUsecase(dbConn, articleRepo, authorRepo, tagsRepo, avRepo, conf, cacheArticleRepo)

	gr := e.Group("/api/v1")
	NewArticleHTTP(gr, usecase, validate)
}

func TestArticleHTTP_GetAll(t *testing.T) {
	SetupTestHTTP(t)

	// Atur query: page=1&limit=10
	req := httptest.NewRequest(http.MethodGet, "/api/v1/articles?page=1&limit=10", nil)
	rec := httptest.NewRecorder()
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), "success"))

}
