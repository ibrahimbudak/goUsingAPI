package main

import (
	"go-api/config"
	"go-api/controller"
	"go-api/repository"
	"go-api/service"
)

var (
	Config *config.Config
)

func main() {
	Config = config.NewConfiguration()

	pg, err := repository.NewPostgres(Config.Postgres).Connect()
	if err != nil {
		panic(err)
	}

	r := repository.NewPersonRepository(pg)
	s := service.NewPersonService(r)
	controller.NewPersonController(s)
}
