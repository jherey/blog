package controllers

import (
	"net/http"

	"github.com/jherey/blog/api/responses"
)

// Home is the home handler function
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Blog API")
}
