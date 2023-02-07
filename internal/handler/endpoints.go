package handler

import (
	"github.com/AnishriM/expenses-diary/internal/services/tag"
	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
	DB     *tag.DBService
}

func NewHandler(db *tag.DBService) *Handler {
	return &Handler{
		Router: mux.NewRouter(),
		DB:     db,
	}
}
func (h *Handler) SetupRoutes() error {
	println("Setting up Routes")
	h.Router.HandleFunc("/api/tag/{id}", h.GetTagByID).Methods("GET")
	h.Router.HandleFunc("/api/tag", h.GetAllTags).Methods("GET")
	h.Router.HandleFunc("/api/tag/{id}", h.UpdateTag).Methods("PUT")
	h.Router.HandleFunc("/api/tag/{id}", h.DeleteTag).Methods("DELETE")
	h.Router.HandleFunc("/api/tag", h.CreateTag).Methods("POST")
	return nil
}
