package service

import (
	"context"
	"url-shortener/internal/domain"
	"url-shortener/internal/repository"
)

type LinksService struct {
	repo repository.Links
}

func (l *LinksService) Create(ctx context.Context, inp CreateLinkInput) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinksService) GetByShortCode(ctx context.Context, shortCode string) (domain.Link, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LinksService) UpdateByShortCode(ctx context.Context, shortCode string, inp UpdateLinkInput) (domain.Link, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LinksService) DeleteByShortCode(ctx context.Context, shortCode string) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinksService) GetStatsByShortCode(ctx context.Context, shortCode string) (domain.LinkStats, error) {
	//TODO implement me
	panic("implement me")
}

func NewLinksService(repo repository.Links) *LinksService {
	return &LinksService{repo: repo}
}
