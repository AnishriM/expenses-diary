package handler

import (
	"github.com/AnishriM/expenses-diary/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
	DB     *service.DBService
}

func NewHandler(db *service.DBService) *Handler {
	return &Handler{
		Router: mux.NewRouter(),
		DB:     db,
	}
}
func (h *Handler) SetupRoutes() error {
	println("Setting up tag routes")
	h.Router.HandleFunc("/api/tag/{id}", h.GetTagByID).Methods("GET")
	h.Router.HandleFunc("/api/tag", h.GetAllTags).Methods("GET")
	h.Router.HandleFunc("/api/tag/{id}", h.UpdateTag).Methods("PUT")
	h.Router.HandleFunc("/api/tag/{id}", h.DeleteTag).Methods("DELETE")
	h.Router.HandleFunc("/api/tag", h.CreateTag).Methods("POST")

	println("Setting up expense routes")
	h.Router.HandleFunc("/api/expense", h.GetAllExpense).Methods("GET")
	h.Router.HandleFunc("/api/expense", h.PostExpenses).Methods("POST")
	return nil
}
