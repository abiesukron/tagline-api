package repository

import (
	"context"
	"database/sql"
	"tagline_api/model/domain"
)

type BeritaRepository interface {
	Keep(ctx context.Context, tx *sql.Tx, berita domain.Berita) domain.Berita
	// Replace(ctx context.Context, tx *sql.Tx, berita domain.Berita) domain.Berita
	// Remove(ctx context.Context, tx *sql.Tx, berita domain.Berita)
	// Only(ctx context.Context, tx *sql.Tx, beritaId int) (domain.Berita, error)
	// All(ctx context.Context, tx *sql.Tx) []domain.Berita
}
