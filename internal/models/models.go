package models

import "time"

type Banner struct {
	Id        int       `json:"id"`
	TagIds    []Tag     `json:"tag_ids"`
	FeatureId int       `json:"feature_id"`
	Content   Content   `json:"content"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Tag struct {
	Id int `json:"id"`
}

type Content struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Url   string `json:"url"`
}

type BannerListFilter struct {
	Limit     uint64 `json:"limit"`
	Offset    uint64 `json:"offset"`
	FeatureId *int   `json:"featureId"`
	TagId     *int   `json:"tagId"`
}

type BannerInfo struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	Url       string    `json:"url"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	FeatureId int       `json:"featureId"`
	TagId     []int     `json:"tagId"`
}

type UserBannerFilter struct {
	FeatureId       int
	TagId           int
	UseLastRevision *bool
}

type BannerUpdateById struct {
	BannerId  int     `json:"banner_id"`
	TagIds    []int   `json:"tag_ids"`
	FeatureId int     `json:"feature_id"`
	Content   Content `json:"content"`
	IsActive  bool    `json:"is_active"`
}

type CreateBannerRequest struct {
	TagIds    []int   `json:"tag_ids"`
	FeatureId int     `json:"feature_id"`
	Content   Content `json:"content"`
	IsActive  bool    `json:"is_active"`
}

type CreateBannerResponse struct {
	BannerId int `json:"banner_id"`
}

type UserBannerKey struct {
	FeatureId int
	TagId     int
}

type BannerByFeatureAndTag struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	Url       string    `json:"url"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	FeatureId int       `json:"featureId"`
	TagId     int       `json:"tagId"`
}
