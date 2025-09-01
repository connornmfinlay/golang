package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sikozonpc/ecom/types"
	"github.com/sikozonpc/ecom/utils"
)

type Handler struct {

}

func NewHandler() *Handler {
  return &Handler{}
}

func ( h *Handler) RegisterRoutes(router *mux.Router) {
  router.HandleFunc("/login", h.handleLogin).Methods("POST")
  router.HandleFunc("/register", h.handleRegister()).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request){

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request){
// get json payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
// check if the user exists



// if not create it
}
