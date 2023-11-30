package main

import (
	"net/http"
	"tagline_api/app"
	"tagline_api/controller"
	"tagline_api/model/repository"
	"tagline_api/model/service"

	"github.com/go-playground/validator"
)

type Postingan struct {
}

func main() {

	DB := app.NewDB()
	validate := validator.New()
	beritaRepository := repository.NewBeritaRepository()
	beritaService := service.NewBeritaService(beritaRepository, DB, validate)
	beritaController := controller.NewBeritaController(beritaService)

	router := app.NewRouter(beritaController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
