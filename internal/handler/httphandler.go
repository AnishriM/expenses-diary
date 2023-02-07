package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AnishriM/expenses-diary/internal/response"
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

func (h *Handler) GetTagByID(w http.ResponseWriter, r *http.Request) {
	println("hit GetTagByID endpoint")
	vars := mux.Vars(r)
	id := vars["id"]
	var uid uint64
	var err error
	if uid, err = strconv.ParseUint(id, 10, 64); err != nil {
		response.SendErrorResponse(w, response.Response{
			Message: "Error ocurred while parsing the int",
			Error:   err,
		})
		return
	}
	tag, err := tag.GetTagByID(h.DB, uint(uid))
	if err != nil {
		response.SendErrorResponse(w, response.Response{
			Message: "Error ocurred getting tag",
			Error:   err,
		})
		return
	}
	response.SendOkResponse(w, tag)
}

func (h *Handler) GetAllTags(w http.ResponseWriter, r *http.Request) {
	println("hit GetAllTags endpoint")
	tags, err := tag.GetAllTags(h.DB)
	if err != nil {
		response.SendErrorResponse(w, response.Response{
			Message: "Error ocurred getting tags",
			Error:   err,
		})
		return
	}
	response.SendOkResponse(w, tags)
}

func (h *Handler) UpdateTag(w http.ResponseWriter, r *http.Request) {
	println("hit UpdateTag endpoint")
	vars := mux.Vars(r)
	id := vars["id"]
	var uid uint64
	var err error

	if uid, err = strconv.ParseUint(id, 10, 64); err != nil {
		response.SendErrorResponse(w, response.Response{
			Message: "Error ocurred getting tags",
			Error:   err,
		})
		return
	}

	var newTag tag.Tag
	if err := json.NewDecoder(r.Body).Decode(&newTag); err != nil {
		response.SendErrorResponse(w, response.Response{
			Message: "Error ocurred decoding body",
			Error:   err,
		})
		return
	}

	tag, err := tag.UpdateTag(uint(uid), newTag.Name, h.DB)
	if err != nil {
		response.SendErrorResponse(w, response.Response{
			Message: "Error ocurred updating tags",
			Error:   err,
		})
		return
	}
	response.SendOkResponse(w, tag)
}

func (h *Handler) DeleteTag(w http.ResponseWriter, r *http.Request) {
	println("hit DeleteTag endpoint")
	vars := mux.Vars(r)
	id := vars["id"]
	var uid uint64
	var err error

	if uid, err = strconv.ParseUint(id, 10, 64); err != nil {
		response.SendErrorResponse(w, response.Response{
			Message: "Error ocurred deleting tag",
			Error:   err,
		})
		return
	}

	tag, err := tag.DeleteTag(uint(uid), h.DB)
	if err != nil {
		response.SendErrorResponse(w, response.Response{
			Message: "Error ocurred updating tags",
			Error:   err,
		})
		return
	}
	response.SendOkResponse(w, tag)
}

func (h *Handler) CreateTag(w http.ResponseWriter, r *http.Request) {
	println("hit CreateTag endpoint")
	var err error
	var newtag tag.Tag

	if err := json.NewDecoder(r.Body).Decode(&newtag); err != nil {
		response.SendErrorResponse(w, response.Response{
			Message: "Error occured whil decoding json body",
			Error:   err,
		})
	}
	println("Tag Name:" + newtag.Name)

	newtag, err = tag.CreateTag(newtag.Name, h.DB)
	if err != nil {
		response.SendErrorResponse(w, response.Response{
			Message: "Error ocurred updating tags",
			Error:   err,
		})
		return
	}
	response.SendOkResponse(w, newtag)
}
