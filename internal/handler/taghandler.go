package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	errmessage "github.com/AnishriM/expenses-diary/internal/common"
	"github.com/AnishriM/expenses-diary/internal/response"
	"github.com/AnishriM/expenses-diary/internal/service"
	"github.com/gorilla/mux"
)

func (h *Handler) GetTagByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var uid uint64
	var err error
	if uid, err = strconv.ParseUint(id, 10, 64); err != nil {
		response.SendErrorResponse(w, response.Response{
			Message: errmessage.GET_TAG_ERROR,
			Error:   err.Error(),
		})
		return
	}
	tag, err := service.GetTagByID(h.DB, uint(uid))
	if err != nil {
		response.SendErrorResponse(w, response.Response{
			Message: errmessage.GET_TAG_ERROR,
			Error:   err.Error(),
		})
		return
	}
	response.SendOkResponse(w, tag)
}

func (h *Handler) GetAllTags(w http.ResponseWriter, r *http.Request) {
	tags, err := service.GetAllTags(h.DB)
	if err != nil {
		response.SendErrorResponse(w, response.Response{
			Message: errmessage.GET_TAG_ERROR,
			Error:   err.Error(),
		})
		return
	}
	response.SendOkResponse(w, tags)
}

func (h *Handler) UpdateTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var uid uint64
	var err error

	if uid, err = strconv.ParseUint(id, 10, 64); err != nil {
		response.SendErrorResponse(w, response.Response{
			Message: errmessage.GET_TAG_ERROR,
			Error:   err.Error(),
		})
		return
	}

	var newTag service.Tag
	if err := json.NewDecoder(r.Body).Decode(&newTag); err != nil {
		response.SendErrorResponse(w, response.Response{
			Message: errmessage.JSON_DECODE_ERROR,
			Error:   err.Error(),
		})
		return
	}

	tag, err := service.UpdateTag(uint(uid), newTag.Name, h.DB)
	if err != nil {
		response.SendErrorResponse(w, response.Response{
			Message: errmessage.UPDATE_TAG_ERROR,
			Error:   err.Error(),
		})
		return
	}
	response.SendOkResponse(w, tag)
}

func (h *Handler) DeleteTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var uid uint64
	var err error

	if uid, err = strconv.ParseUint(id, 10, 64); err != nil {
		response.SendErrorResponse(w, response.Response{
			Message: errmessage.DELETE_TAG_ERROR,
			Error:   err.Error(),
		})
		return
	}

	tag, err := service.DeleteTag(uint(uid), h.DB)
	if err != nil {
		response.SendErrorResponse(w, response.Response{
			Message: errmessage.DELETE_TAG_ERROR,
			Error:   err.Error(),
		})
		return
	}
	response.SendOkResponse(w, tag)
}

func (h *Handler) CreateTag(w http.ResponseWriter, r *http.Request) {
	var err error
	var newtag service.Tag

	if err := json.NewDecoder(r.Body).Decode(&newtag); err != nil {
		response.SendErrorResponse(w, response.Response{
			Message: errmessage.JSON_DECODE_ERROR,
			Error:   err.Error(),
		})
	}

	newtag, err = service.CreateTag(newtag.Name, h.DB)
	if err != nil {
		response.SendErrorResponse(w, response.Response{
			Message: errmessage.CREATE_TAG_ERROR,
			Error:   err.Error(),
		})
		return
	}
	response.SendOkResponse(w, newtag)
}
