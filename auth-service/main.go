package main

import (
	"auth-service/internal/domain/user/repository"
	"auth-service/internal/domain/user/usecase"
	"auth-service/package/config"
	"auth-service/package/connection/database"
	"os"

	api "auth-service/cmd/api"
	"auth-service/cmd/migrate"

	"github.com/urfave/cli/v2"
)

func main() {
	dbConf := config.NewDatabase()
	conf := config.NewConfig()
	conn := database.WebDB
	dbConn := database.NewDatabase(conn, dbConf)

	repo := repository.NewUserRepository(dbConn)
	usecase := usecase.NewUserUsecase(repo, conf)
	
	cmds := []*cli.Command{}
	cmds = append(cmds, api.ServeAPI(usecase, conf)...)
	cmds = append(cmds, migrate.NewMigrate(dbConn)...)

	app := &cli.App{
		Name:     "auth-service",
		Commands: cmds,
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
