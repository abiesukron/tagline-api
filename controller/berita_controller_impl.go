package controller

import (
	"net/http"
	"tagline_api/helper"
	"tagline_api/model/service"
	"tagline_api/model/web"

	"github.com/julienschmidt/httprouter"
)

type BeritaControllerImpl struct {
	BeritaService service.BeritaService
}

func NewBeritaController(beritaService service.BeritaService) BeritaContorller {
	return &BeritaControllerImpl{
		BeritaService: beritaService,
	}
}

func (controller *BeritaControllerImpl) Keep(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	beritaCreateRequest := web.BeritaCreateRequest{}
	helper.ReadFromRequestBody(request, &beritaCreateRequest)
	beritaResponse := controller.BeritaService.Keep(request.Context(), beritaCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   beritaResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
