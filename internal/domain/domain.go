package domain

import "time"

type Banner struct {
	Id        int
	Title     string
	Text      string
	Url       string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	FeatureId int
}

type Tag struct {
	Id       int
	BannerId int
}

type Feature struct {
	Id int
}

type Role struct {
	UserId string
	Role   string
}
