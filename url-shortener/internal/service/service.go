package service

import (
	"context"
	"url-shortener/internal/domain"
	"url-shortener/internal/repository"
)

type Links interface {
	Create(ctx context.Context, inp CreateLinkInput) error
	GetByShortCode(ctx context.Context, shortCode string) (domain.Link, error)
	UpdateByShortCode(ctx context.Context, shortCode string, inp UpdateLinkInput) (domain.Link, error)
	DeleteByShortCode(ctx context.Context, shortCode string) error
	GetStatsByShortCode(ctx context.Context, shortCode string) (domain.LinkStats, error)
}

type CreateLinkInput struct {
	URL string
}

type UpdateLinkInput struct {
	URL string
}

type Services struct {
	Links Links
}

type Dependencies struct {
	Repos *repository.Repositories
}

func NewServices(dependencies *repository.Repositories) *Services {
	linksService := NewLinksService(dependencies.Links)
	return &Services{Links: linksService}
}
