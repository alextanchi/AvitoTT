package service

import (
	"AvitoTestTask/internal/convert"
	"AvitoTestTask/internal/domain"
	"AvitoTestTask/internal/models"
	"AvitoTestTask/internal/repository"
	"context"
)

type Service interface {
	CreateBanner(ctx context.Context, banner domain.Banner) (int, error)
	DeleteBannerById(ctx context.Context, id int) error
	GetBanner(ctx context.Context, input models.UserBannerFilter) (domain.Banner, error)
	GetBannersByFeatureAndTag(ctx context.Context, filter models.BannerListFilter) ([]models.BannerByFeatureAndTag, error)
	UpdateBannerById(ctx context.Context, banner domain.Banner, request models.BannerUpdateById) error
}
type BannerService struct {
	storage repository.Repository
}

func NewService(repos repository.Repository) Service {

	return BannerService{
		storage: repos,
	}
}

func (b BannerService) CreateBanner(ctx context.Context, banner domain.Banner) (int, error) {
	return b.storage.CreateBanner(ctx, banner)

}

func (b BannerService) DeleteBannerById(ctx context.Context, id int) error {
	return b.storage.DeleteBannerById(ctx, id)

}

func (b BannerService) GetBanner(ctx context.Context, input models.UserBannerFilter) (domain.Banner, error) {

	//поменять параметры и обработать if *input.UseLastRevision { когда флаг прилетает, что делаем?

	b.storage.GetBanner()

}

func (b BannerService) GetBannersByFeatureAndTag(ctx context.Context, filter models.BannerListFilter) ([]models.BannerByFeatureAndTag, error) {
	result, err := b.storage.GetBannersByFeatureAndTag(ctx, filter)
	if err != nil {
		return []models.BannerByFeatureAndTag{}, err
	}

	return result, nil

}

func (b BannerService) UpdateBannerById(ctx context.Context, request models.BannerUpdateById) error {
	err := b.storage.DeleteTagByBannerId(ctx, request.BannerId)
	if err != nil {
		return err
	}
	err = b.storage.SetTags(ctx, convert.BannerToTag(request))
	if err != nil {
		return err
	}
	err = b.storage.UpdateBannerById(ctx, request)
	if err != nil {
		return err
	}

	return nil
}
