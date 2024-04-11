package repository

import (
	"AvitoTestTask/internal/domain"
	"AvitoTestTask/internal/models"
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"log"
)

type Repository interface {
	CreateBanner(ctx context.Context, banner domain.Banner) (int, error)
	DeleteBannerById(ctx context.Context, id int) error
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
		Insert(tableBanner).Columns(
		"id",
		"title",
		"text",
		"url",
		"is_active",
		"created_at",
		"updated_at",
		"feature_id",
	).
		Values(
			banner.Id,
			banner.Title,
			banner.Text,
			banner.Url,
			banner.IsActive,
			banner.CreatedAt,
			banner.UpdatedAt,
			banner.FeatureId,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return 0, err
	}
	_, err = b.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return banner.Id, nil

}

func (b Banner) DeleteBannerById(ctx context.Context, id int) error {

	deleteQuery := sq.Delete("banner").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := deleteQuery.ToSql()
	if err != nil {
		return err
	}
	_, err = b.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (b Banner) GetBanner(ctx context.Context, input models.UserBannerFilter) (domain.Banner, error) {

}

func (b Banner) GetBannersByFeatureAndTag(ctx context.Context, filter models.BannerListFilter) ([]models.BannerByFeatureAndTag, error) {
	result := make([]models.BannerByFeatureAndTag, 0)
	queryBuilder := sq.
		Select(
			"b.id",
			"b.title",
			"b.text",
			"b.url",
			"b.is_active",
			"b.created_at",
			"b.updated_at",
			"f.id",
			"t.id",
		).
		From("banner b").
		Join("feature f ON b.feature_id = f.id").
		Join("tag t ON b.id = t.banner_id").
		OrderBy("banner.id").
		Limit(filter.Limit).
		Offset(filter.Offset).
		PlaceholderFormat(sq.Dollar)

	if filter.FeatureId != nil {
		queryBuilder = queryBuilder.Where(sq.Eq{"feature.id": filter.FeatureId})
	}
	if filter.TagId != nil {
		queryBuilder = queryBuilder.Where(sq.Eq{"feature.id": filter.TagId})
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := b.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var banner models.BannerByFeatureAndTag
		err := rows.Scan(
			&banner.Id,
			&banner.Title,
			&banner.Text,
			&banner.Url,
			&banner.IsActive,
			&banner.CreatedAt,
			&banner.UpdatedAt,
			&banner.FeatureId,
			&banner.TagId,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, banner)
	}

	return result, nil
}

func (b Banner) UpdateBannerById(ctx context.Context, banner domain.Banner) error {

	updateQuery := sq.Update("banner").
		Set("title", title).
		Set("text", text).
		Set("url", url).
		Set("is_active", isActive).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := udpateQuery.ToSql()
	if err != nil {
		return err
	}
	_, err = b.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
