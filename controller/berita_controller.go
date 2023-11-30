package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BeritaContorller interface {
	Keep(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	// Replace()
	// Remove()
	// Only()
	// All()
}
