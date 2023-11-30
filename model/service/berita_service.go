package service

import (
	"context"
	"tagline_api/model/web"
)

type BeritaService interface {
	Keep(ctx context.Context, request web.BeritaCreateRequest) web.BeritaResponse
	// Replace(ctx context.Context, request web.BeritaUpdateRequest) web.BeritaResponse
	// Remove(ctx context.Context, beritaId int)
	// Only(ctx context.Context, beritaId int) web.BeritaResponse
	// All(ctx context.Context) []web.BeritaResponse
}
