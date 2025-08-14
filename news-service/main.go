package main

import (
	"os"

	avRepository "news-service/internal/domain/article-version/repository"
	"news-service/internal/domain/article/repository"
	articleUsecase "news-service/internal/domain/article/usecase"

	avUsecase "news-service/internal/domain/article-version/usecase"
	authorRepository "news-service/internal/domain/author/repository"
	tagsRepository "news-service/internal/domain/tags/repository"

	"news-service/package/config"
	"news-service/package/connection/cache"
	database "news-service/package/connection/database"

	api "news-service/cmd/api"
	"news-service/cmd/migrate"

	"github.com/urfave/cli/v2"

	cacheRepo "news-service/internal/domain/article/repository/cache"

	"news-service/package/validator"
)

func main() {

  dbConf := config.NewDatabase()
  conf := config.NewConfig()
  dbConn := database.NewDatabase("wdb",dbConf)
  cacheConf := config.NewCache()
  cache := cache.NewRedis("cache-cms", cacheConf)

  articleRepo := repository.NewsRepository(dbConn, cache)
  authorRepo := authorRepository.NewAuthorRepository(dbConn)
  tagsRepo := tagsRepository.NewTagsRepository(dbConn)
  avRepo := avRepository.NewArticleVersionRepository(dbConn)
  cacheArticleRepo := cacheRepo.NewCacheRepository(cache)
  validate := validator.NewValidator()

  artUsecase := articleUsecase.NewsUsecase(dbConn, articleRepo, authorRepo, tagsRepo, avRepo, conf, cacheArticleRepo)
  avUc := avUsecase.NewArticleVersionUsecase(avRepo, tagsRepo)
  cmds := []*cli.Command{}
  cmds = append(cmds, api.ServeAPI(conf, validate, artUsecase, avUc)...)
  cmds = append(cmds, migrate.NewMigrate(dbConn)...)
  app := &cli.App{
    Name: "news-service",
    Commands: cmds,
  }

  if err := app.Run(os.Args); err != nil {
    panic(err)
  }
}
