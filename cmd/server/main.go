package main

import (
	"net/http"

	"github.com/AnishriM/expenses-diary/internal/handler"
	"github.com/AnishriM/expenses-diary/internal/service"
)

type Application struct {
	Handler *handler.Handler
}

func (application *Application) Run() error {
	db, err := service.NewDatabase()
	if err != nil {
		println("Error occurred while setting up database.")
		return err
	}

	if err = service.MigrateDB(db); err != nil {
		return err
	}

	println("Setting up application")
	application.Handler = handler.NewHandler(service.NewService(db))
	application.Handler.SetupRoutes()
	println("Listen and serve")
	http.ListenAndServe(":8080", application.Handler.Router)
	return nil
}

func main() {
	application := Application{}
	application.Run()
}
