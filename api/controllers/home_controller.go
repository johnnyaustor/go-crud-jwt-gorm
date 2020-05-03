package controllers

import (
	"net/http"

	"github.com/johnnyaustor/go-crud-jwt/api/responses"
)

// Home Controller
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to this API")
}
