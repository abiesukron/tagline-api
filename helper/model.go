package helper

import (
	"tagline_api/model/domain"
	"tagline_api/model/web"
)

func ToBeritaResponse(berita domain.Berita) web.BeritaResponse {
	return web.BeritaResponse{
		Id:        berita.Id,
		Judul:     berita.Judul,
		Slug:      berita.Slug,
		Excerpt:   berita.Excerpt,
		Isi:       berita.Isi,
		Tag:       berita.Tag,
		Gambar:    berita.Gambar,
		Kategori:  berita.Kategori,
		Publikasi: berita.Publikasi,
		User:      berita.User,
	}
}

func ToBeritaResponses(semuaberita []domain.Berita) []web.BeritaResponse {
	var beritaResponses []web.BeritaResponse
	for _, berita := range semuaberita {
		beritaResponses = append(beritaResponses, ToBeritaResponse(berita))
	}

	return beritaResponses
}
