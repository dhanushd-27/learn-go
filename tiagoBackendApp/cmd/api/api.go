package api

import (
	"log"
	"net/http"
	"github.com/dhanushd-27/learn-go/tiagoBackendApp/services/user"
	"github.com/gorilla/mux"
)


type APIServer struct {
	addr string
	// db *sql.DB
}

func NewApiServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)
 
	return http.ListenAndServe(s.addr, router)
}