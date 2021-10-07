package main

import (
	"log"

	"github.com/K0STYAa/AvitoTech"
	"github.com/K0STYAa/AvitoTech/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(AvitoTech.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}