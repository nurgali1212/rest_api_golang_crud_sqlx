package main

import (
	"test_project/database"
	"test_project/handler"
	"test_project/repository"
	"test_project/service"

	_ "github.com/lib/pq"
)

func main() {
	db := database.InitDatabase()
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	router := handlers.InitRoutes()

	router.Run()
}
