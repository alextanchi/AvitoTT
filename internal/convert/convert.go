package convert

import (
	"AvitoTestTask/internal/domain"
	"AvitoTestTask/internal/models"
	"log"
	"time"
)

func BannerToTag(banner models.BannerUpdateById) []domain.Tag {
	result := make([]domain.Tag, len(banner.TagIds))

	for i, el := range banner.TagIds {
		tag := domain.Tag{
			Id:       el,
			BannerId: banner.BannerId,
		}
		result[i] = tag
	}
	return result
}

func DomainBannerToContent(banner domain.Banner) models.Content {
	return models.Content{
		Title: banner.Title,
		Text:  banner.Text,
		Url:   banner.Url,
	}

}

func BannerByFeatureAndTagToCache(input []models.BannerByFeatureAndTag) map[models.UserBannerKey]domain.Banner {
	result := make(map[models.UserBannerKey]domain.Banner)
	for _, el := range input {
		val := domain.Banner{
			Id:        el.Id,
			Title:     el.Title,
			Text:      el.Text,
			Url:       el.Url,
			IsActive:  el.IsActive,
			CreatedAt: el.CreatedAt,
			UpdatedAt: el.UpdatedAt,
			FeatureId: el.FeatureId,
		}
		key := models.UserBannerKey{
			FeatureId: el.FeatureId,
			TagId:     el.TagId,
		}

		if _, ok := result[key]; ok {
			continue
		}
		result[key] = val

	}
	return result
}

func CreateBannerRequestBannerDomainToModel(banner models.CreateBannerRequest) domain.Banner {
	now := time.Now()
	return domain.Banner{
		Title:     banner.Content.Title,
		Text:      banner.Content.Text,
		Url:       banner.Content.Url,
		IsActive:  banner.IsActive,
		CreatedAt: now,
		UpdatedAt: now,
		FeatureId: banner.FeatureId,
	}
}
func BannerByFeatureAndTagToBanner(banner []models.BannerByFeatureAndTag) []models.Banner {
	log.Println(banner)
	result := make([]models.Banner, 0, len(banner))
	m := make(map[int]models.Banner)
	for _, el := range banner {
		if _, ok := m[el.Id]; ok {
			mod := m[el.Id]
			mod.TagIds = append(mod.TagIds, models.Tag{Id: el.TagId})
			continue
		}
		m[el.Id] = models.Banner{
			Id:        el.Id,
			FeatureId: el.FeatureId,
			IsActive:  el.IsActive,
			CreatedAt: el.CreatedAt,
			UpdatedAt: el.UpdatedAt,
			TagIds:    []models.Tag{{Id: el.TagId}},
			Content: models.Content{
				Title: el.Title,
				Text:  el.Text,
				Url:   el.Url,
			},
		}
	}
	for _, el := range m {
		result = append(result, el)
	}
	return result
}
