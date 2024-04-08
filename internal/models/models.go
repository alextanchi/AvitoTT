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
