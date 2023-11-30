package repository

import (
	"context"
	"database/sql"
	"tagline_api/helper"
	"tagline_api/model/domain"
)

type BeritaRepositoryImpl struct{}

func NewBeritaRepository() BeritaRepository {
	return &BeritaRepositoryImpl{}
}

func (repository *BeritaRepositoryImpl) Keep(ctx context.Context, tx *sql.Tx, berita domain.Berita) domain.Berita {
	SQL := "INSERT INTO berita(judul, slug, excerpt, isi, tag, gambar, kategoriid, publikasi, usersid) VALUES(?,?,?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, berita.Judul, berita.Slug, berita.Excerpt, berita.Isi, berita.Tag, berita.Gambar, berita.Kategori, berita.Publikasi, berita.User)
	helper.PanicErrorHandle(err)

	id, err := result.LastInsertId()
	helper.PanicErrorHandle(err)

	berita.Id = int(id)
	return berita
}

// func (repository *BeritaRepositoryImpl) Replace(ctx context.Context, tx *sql.DB, berita domain.Berita) domain.Berita {

// }

// func (repository *BeritaRepositoryImpl) Remove(ctx context.Context, tx *sql.DB, berita domain.Berita) {

// }

// func (repository *BeritaRepositoryImpl) Only(ctx context.Context, tx *sql.DB, beritaId int) (domain.Berita, error) {

// }

// func (repository *BeritaRepositoryImpl) All(ctx context.Context, tx *sql.DB) []domain.Berita {

// }
