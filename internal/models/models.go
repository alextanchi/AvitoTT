package models

import "time"

type Banner struct {
	Id        int       `json:"id"`
	TagIds    []Tag     `json:"tag_ids"`
	FeatureId int       `json:"feature_id"`
	Content   Content   `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
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
	Limit     uint64
	Offset    uint64
	FeatureId *int
	TagId     *int
}

type BannerByFeatureAndTag struct {
	Id        int
	Title     string
	Text      string
	Url       string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	FeatureId int
	TagId     int
}

type UserBannerFilter struct {
	FeatureId       int
	TagId           int
	UseLastRevision *bool
}
type UserBannerResponse struct {
	Title string
	Text  string
	Url   string
}
