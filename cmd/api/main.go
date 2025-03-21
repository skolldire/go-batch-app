package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/skolldire/go-web-app/cmd/api/routes/ping"
	_ "github.com/skolldire/go-web-app/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", ping.NewService().Apply())
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
