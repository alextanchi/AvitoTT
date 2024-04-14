package service

import (
	"AvitoTestTask/internal/cache"
	"AvitoTestTask/internal/convert"
	"AvitoTestTask/internal/domain"
	"AvitoTestTask/internal/models"
	"AvitoTestTask/internal/repository"
	"context"
)

type Service interface {
	CreateBanner(ctx context.Context, banner models.CreateBannerRequest) (int, error)
	DeleteBannerById(ctx context.Context, id int) error
	GetBanner(ctx context.Context, input models.UserBannerFilter) (domain.Banner, error)
	GetBannersByFeatureAndTag(ctx context.Context, filter models.BannerListFilter) ([]models.BannerByFeatureAndTag, error)
	UpdateBannerById(ctx context.Context, request models.BannerUpdateById) error
}
type BannerService struct {
	storage repository.Repository
	cache   cache.IBannerCache
}

func NewService(repos repository.Repository, bannerCache cache.IBannerCache) Service {
	return BannerService{
		storage: repos,
		cache:   bannerCache,
	}
}

func (b BannerService) CreateBanner(ctx context.Context, banner models.CreateBannerRequest) (int, error) {
	id, err := b.storage.CreateBanner(ctx, convert.CreateBannerRequestBannerDomainToModel(banner))
	if err != nil {
		return 0, err
	}
	err = b.storage.SetTags(ctx, convert.BannerToTag(models.BannerUpdateById{
		TagIds:   banner.TagIds,
		BannerId: id,
	}))
	if err != nil {
		return 0, err
	}
	return id, nil

}

func (b BannerService) DeleteBannerById(ctx context.Context, id int) error {
	return b.storage.DeleteBannerById(ctx, id)
}

func (b BannerService) GetBanner(ctx context.Context, input models.UserBannerFilter) (domain.Banner, error) {
	if input.UseLastRevision != nil {
		if *input.UseLastRevision == true {
			result, err := b.storage.GetBanner(ctx, input)
			if err != nil {
				return domain.Banner{}, err
			}
			return result, nil
		}
	}
	return b.cache.Get(models.UserBannerKey{
		FeatureId: input.FeatureId,
		TagId:     input.TagId,
	})

}

func (b BannerService) GetBannersByFeatureAndTag(ctx context.Context, filter models.BannerListFilter) ([]models.BannerByFeatureAndTag, error) {
	return b.storage.GetBannersByFeatureAndTag(ctx, filter)
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
