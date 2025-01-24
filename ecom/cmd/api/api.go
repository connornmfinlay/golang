package main

import (
	"database/sql"
	"log"
	"net/http"

	user "github.com/connornmfinlay/golang/ecom/cmd/service-user"
	"github.com/gorilla/mux"
)


type APIServer struct { 
  addr string
  db *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
  return &APIServer{
    addr: addr,
    db: db,
  }
}

func (s *APIServer) Run() error { //all routing done here
  router := mux.NewRouter()
  subrouter := router.PathPrefix("/api/v1").Subrouter()
  
  userHandler := user.NewHandler()
  userHandler .RegisterRoutes(subrouter)

  log.Println("Listening on", s.addr)
  
  return http.ListenAndServe(s.addr, router) 
}
