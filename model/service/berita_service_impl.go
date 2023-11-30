package service

import (
	"context"
	"database/sql"
	"tagline_api/helper"
	"tagline_api/model/domain"
	"tagline_api/model/repository"
	"tagline_api/model/web"

	"github.com/go-playground/validator"
)

type BeritaServiceImpl struct {
	BeritaRepository repository.BeritaRepository
	DB               *sql.DB
	Validator        *validator.Validate
}

func NewBeritaService(beritaRepository repository.BeritaRepository, DB *sql.DB, validator *validator.Validate) BeritaService {
	return &BeritaServiceImpl{
		BeritaRepository: beritaRepository,
		DB:               DB,
		Validator:        validator,
	}
}

func (service *BeritaServiceImpl) Keep(ctx context.Context, request web.BeritaCreateRequest) web.BeritaResponse {
	err := service.Validator.Struct(request)
	helper.PanicErrorHandle(err)
	tx, err := service.DB.Begin()
	helper.PanicErrorHandle(err)
	defer service.DB.Close()
	berita := domain.Berita{
		Judul:     request.Judul,
		Slug:      request.Slug,
		Excerpt:   request.Excerpt,
		Isi:       request.Isi,
		Tag:       request.Tag,
		Gambar:    request.Gambar,
		Kategori:  request.Kategori,
		Publikasi: request.Publikasi,
		User:      request.User,
	}

	berita = service.BeritaRepository.Keep(ctx, tx, berita)

	return helper.ToBeritaResponse(berita)
}
