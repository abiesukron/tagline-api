package app

import (
	"tagline_api/controller"
	"tagline_api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(beritaController controller.BeritaContorller) *httprouter.Router {
	router := httprouter.New()
	router.POST("/api/berita", beritaController.Keep)
	router.PanicHandler = exception.ErrorHandler
	return router
}
