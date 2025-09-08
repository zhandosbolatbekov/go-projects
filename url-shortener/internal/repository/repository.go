package repository

import (
	"context"
	"url-shortener/internal/domain"
	mock_repository "url-shortener/internal/repository/mocks"
)

type Links interface {
	Create(ctx context.Context, link domain.Link) error
	GetByShortCode(ctx context.Context, shortCode string) (domain.Link, error)
	UpdateByShortCode(ctx context.Context, shortCode string, newURL string) (domain.Link, error)
	DeleteByShortCode(ctx context.Context, shortCode string) error
	GetStatsByShortCode(ctx context.Context, shortCode string) (domain.LinkStats, error)
}

type Repositories struct {
	Links Links
}

func NewRepositories() *Repositories {
	return &Repositories{
		Links: mock_repository.NewMockLinks(),
	}
}
