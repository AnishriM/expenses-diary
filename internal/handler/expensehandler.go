package handler

import (
	"net/http"

	"github.com/AnishriM/expenses-diary/internal/response"
	"github.com/AnishriM/expenses-diary/internal/service"
)

func (h *Handler) GetAllExpense(w http.ResponseWriter, r *http.Request) {
	expense, err := service.GetAllExpenses(h.DB)
	if err != nil {
		response.SendErrorResponse(w, response.Response{
			Message: "Error occurred while fetching expenses information",
			Error:   err.Error(),
		})
	}
	response.SendOkResponse(w, expense)
}

func (h *Handler) PostExpenses(w http.ResponseWriter, r *http.Request) {
	expense, err := service.PostExpenses(h.DB)
	if err != nil {
		response.SendErrorResponse(w, response.Response{
			Message: "Error occurred while inserting expense data",
			Error:   err.Error(),
		})
	}
	response.SendOkResponse(w, expense)
}
