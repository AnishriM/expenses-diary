package main

import (
	"net/http"

	"github.com/AnishriM/expenses-diary/internal/handler"
	database "github.com/AnishriM/expenses-diary/internal/services/tag"
)

type Application struct {
	Handler *handler.Handler
}

func (application *Application) Run() error {
	db, err := database.NewDatabase()
	if err != nil {
		println("Error occurred while setting up database.")
		return err
	}

	if err = database.MigrateDB(db); err != nil {
		return err
	}

	println("Setting up application")
	application.Handler = handler.NewHandler(database.NewService(db))
	application.Handler.SetupRoutes()
	println("Listen and serve")
	http.ListenAndServe(":8080", application.Handler.Router)
	return nil
}

func main() {
	application := Application{}
	application.Run()
}
