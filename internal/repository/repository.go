package repository

import (
	"AvitoTestTask/internal/domain"
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"log"
)

type Repository interface {
	CreateBanner(ctx context.Context, banner domain.Banner) (int, error)
}

type Banner struct {
	db *sql.DB
}

func NewBanner(db *sql.DB) Repository { //конструктор

	return &Banner{
		db: db,
	}
}

// ConnectDb добавляем подключение к базе
func ConnectDb() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "password", "banners")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	log.Println("подключились к БД")

	return db, nil
}

func (b Banner) CreateBanner(ctx context.Context, banner domain.Banner) (int, error) {

	query, args, err := sq.
		Insert(tableBanner).Columns("id", "title", "text", "url", "is_active", "created_at", "updated_at", "feature_id").
		Values(banner.Id, banner.Title, banner.Text, banner.Url, banner.IsActive, banner.CreatedAt, banner.UpdatedAt, banner.FeatureId).
		ToSql()
	if err != nil {
		return 0, err
	}
	b.db.ExecContext(ctx, query, args...)

}
