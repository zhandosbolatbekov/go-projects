package mock_repository

import (
	"context"
	"errors"
	"time"
	"url-shortener/internal/domain"
)

type MockLinks struct {
	dict         map[int64]domain.Link
	shortCodeIdx map[string]int64
	stats        map[int64]int
}

func (m *MockLinks) Create(ctx context.Context, link domain.Link) error {
	_, exist := m.dict[link.ID]
	if exist {
		return errors.New("link already exists")
	}
	m.dict[link.ID] = link
	m.shortCodeIdx[link.ShortCode] = link.ID
	return nil
}

func (m *MockLinks) GetByShortCode(ctx context.Context, shortCode string) (domain.Link, error) {
	id, exist := m.shortCodeIdx[shortCode]
	if !exist {
		return domain.Link{}, errors.New("link not found")
	}
	return m.dict[id], nil
}

func (m *MockLinks) UpdateByShortCode(ctx context.Context, shortCode string, newURL string) (domain.Link, error) {
	id, exist := m.shortCodeIdx[shortCode]
	if !exist {
		return domain.Link{}, errors.New("link not found")
	}
	link := m.dict[id]
	newLink := domain.NewLink(
		link.ID,
		newURL,
		link.ShortCode,
		link.CreatedAt,
		time.Now(),
	)
	m.dict[id] = *newLink
	return *newLink, nil
}

func (m *MockLinks) DeleteByShortCode(ctx context.Context, shortCode string) error {
	id, exist := m.shortCodeIdx[shortCode]
	if !exist {
		return errors.New("link not found")
	}
	delete(m.dict, id)
	return nil
}

func (m *MockLinks) GetStatsByShortCode(ctx context.Context, shortCode string) (domain.LinkStats, error) {
	id, exist := m.shortCodeIdx[shortCode]
	if !exist {
		return domain.LinkStats{}, errors.New("link not found")
	}
	stats := domain.NewLinkStats(m.dict[id], m.stats[id])
	return *stats, nil
}

func NewMockLinks() *MockLinks {
	return &MockLinks{}
}
