package cache

import (
	"AvitoTestTask/internal/convert"
	"AvitoTestTask/internal/domain"
	"AvitoTestTask/internal/models"
	"AvitoTestTask/internal/repository"
	"context"
	"errors"
	"sync"
)

type IBannerCache interface {
	Refresh(ctx context.Context) error
	Get(key models.UserBannerKey) (domain.Banner, error)
}
type BannerCache struct {
	mu   *sync.Mutex
	m    map[models.UserBannerKey]domain.Banner
	repo repository.Repository
}

func NewBannerCache(repo repository.Repository) IBannerCache {

	return &BannerCache{
		mu:   &sync.Mutex{},
		m:    make(map[models.UserBannerKey]domain.Banner),
		repo: repo,
	}
}

func (b *BannerCache) Refresh(ctx context.Context) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	result, err := b.repo.GetAllBanners(ctx)
	if err != nil {
		return err
	}
	b.m = convert.BannerByFeatureAndTagToCache(result)
	return nil
}
func (b *BannerCache) Get(key models.UserBannerKey) (domain.Banner, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if _, ok := b.m[key]; !ok {
		return domain.Banner{}, errors.New("Не найден баннер по фиче и тэгу")
	}
	return b.m[key], nil
}
