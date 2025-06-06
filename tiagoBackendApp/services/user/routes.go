package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {

}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/user", h.GetUser).Methods("GET")
	router.HandleFunc("/user", h.CreateUser).Methods("POST")
	router.HandleFunc("/user", h.DeleteUser).Methods("DELETE")
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User"))
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create User"))
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}