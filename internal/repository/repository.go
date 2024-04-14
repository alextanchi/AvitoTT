package repository

import (
	"AvitoTestTask/internal/domain"
	"AvitoTestTask/internal/models"
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"log"
)

type Repository interface {
	CreateBanner(ctx context.Context, banner domain.Banner) (int, error)
	DeleteBannerById(ctx context.Context, id int) error
	GetBanner(ctx context.Context, input models.UserBannerFilter) (domain.Banner, error)
	GetBannersByFeatureAndTag(ctx context.Context, filter models.BannerListFilter) ([]models.BannerByFeatureAndTag, error)
	UpdateBannerById(ctx context.Context, banner models.BannerUpdateById) error

	DeleteTagByBannerId(ctx context.Context, id int) error
	SetTags(ctx context.Context, tag []domain.Tag) error
	GetAllBanners(ctx context.Context) ([]models.BannerByFeatureAndTag, error)
	GetRole(ctx context.Context, userId string) (string, error)
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
	query := sq.
		Insert(tableBanner).Columns(
		"title",
		"text",
		"url",
		"is_active",
		"created_at",
		"updated_at",
		"feature_id",
	).
		Values(
			banner.Title,
			banner.Text,
			banner.Url,
			banner.IsActive,
			banner.CreatedAt,
			banner.UpdatedAt,
			banner.FeatureId,
		).
		Suffix("RETURNING \"id\"").
		RunWith(b.db).
		PlaceholderFormat(sq.Dollar)
	var id int
	err := query.QueryRow().Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (b Banner) DeleteBannerById(ctx context.Context, id int) error {
	deleteQuery := sq.Delete("banner").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)
	query, args, err := deleteQuery.ToSql()
	if err != nil {
		return err
	}
	res, err := b.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if rowsAffected == 0 {
		return ErrBannerNotFound
	}
	return nil
}

func (b Banner) GetBanner(ctx context.Context, input models.UserBannerFilter) (domain.Banner, error) {
	selectQuery := sq.Select(
		"b.id",
		"b.title",
		"b.text",
		"b.url",
		"b.is_active",
		"b.created_at",
		"b.updated_at",
		"f.id",
	).
		From("banner b").
		Join("feature f ON b.feature_id = f.id").
		Join("tag t ON b.id = t.banner_id").
		Where(sq.Eq{"f.id": input.FeatureId}).
		Where(sq.Eq{"t.id": input.TagId}).
		PlaceholderFormat(sq.Dollar)
	query, args, err := selectQuery.ToSql()
	if err != nil {
		return domain.Banner{}, err
	}
	rows, err := b.db.QueryContext(ctx, query, args...)
	if err != nil {
		return domain.Banner{}, err
	}
	var banner domain.Banner
	for rows.Next() {
		err := rows.Scan(
			&banner.Id,
			&banner.Title,
			&banner.Text,
			&banner.Url,
			&banner.IsActive,
			&banner.CreatedAt,
			&banner.UpdatedAt,
			&banner.FeatureId,
		)
		if err == sql.ErrNoRows {
			return domain.Banner{}, ErrBannerNotFound
		}
		if err != nil {
			return domain.Banner{}, err
		}
	}
	return banner, nil
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
		RightJoin("tag t ON b.id = t.banner_id").
		OrderBy("b.id").
		Limit(filter.Limit).
		Offset(filter.Offset).
		PlaceholderFormat(sq.Dollar)
	if filter.FeatureId != nil {
		queryBuilder = queryBuilder.Where(sq.Eq{"f.id": filter.FeatureId})
	}
	if filter.TagId != nil {
		queryBuilder = queryBuilder.Where(sq.Eq{"t.id": filter.TagId})
	}
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := b.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
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
		if err == sql.ErrNoRows {
			return []models.BannerByFeatureAndTag{}, nil
		}
		if err != nil {
			return nil, err
		}
		result = append(result, banner)
	}
	return result, nil
}

func (b Banner) UpdateBannerById(ctx context.Context, banner models.BannerUpdateById) error {
	updateQuery := sq.Update("banner").
		Set("title", banner.Content.Title).
		Set("text", banner.Content.Text).
		Set("url", banner.Content.Url).
		Set("is_active", banner.IsActive).
		Set("feature_id", banner.FeatureId).
		Where(sq.Eq{"id": banner.BannerId}).
		PlaceholderFormat(sq.Dollar)
	query, args, err := updateQuery.ToSql()
	if err != nil {
		return err
	}
	result, err := b.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrBannerNotFound
	}
	return nil
}

func (b Banner) DeleteTagByBannerId(ctx context.Context, id int) error {
	deleteQuery := sq.Delete("tag").
		Where(sq.Eq{"banner_id": id}).
		PlaceholderFormat(sq.Dollar)
	query, args, err := deleteQuery.ToSql()
	if err != nil {
		return err
	}
	res, err := b.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if rowsAffected == 0 {
		return ErrBannerNotFound
	}
	return nil
}

func (b Banner) SetTags(ctx context.Context, tag []domain.Tag) error {
	for _, el := range tag {
		query, args, err := sq.
			Insert(tableTag).Columns(
			"id",
			"banner_id",
		).
			Values(
				el.Id,
				el.BannerId,
			).
			PlaceholderFormat(sq.Dollar).
			ToSql()
		if err != nil {
			return err
		}
		_, err = b.db.ExecContext(ctx, query, args...)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b Banner) GetAllBanners(ctx context.Context) ([]models.BannerByFeatureAndTag, error) {
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
		PlaceholderFormat(sq.Dollar)
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := b.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
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
		if err == sql.ErrNoRows {
			return []models.BannerByFeatureAndTag{}, nil
		}
		if err != nil {
			return nil, err
		}
		result = append(result, banner)
	}
	return result, nil
}

func (b Banner) GetRole(ctx context.Context, userId string) (string, error) {
	selectQuery := sq.Select(
		"role",
	).
		From(tableRole).
		Where(sq.Eq{"user_id": userId}).
		PlaceholderFormat(sq.Dollar)
	query, args, err := selectQuery.ToSql()
	if err != nil {
		return "", err
	}
	rows, err := b.db.QueryContext(ctx, query, args...)
	if err != nil {
		return "", err
	}
	var role string
	for rows.Next() {
		err := rows.Scan(
			&role,
		)
		if err != nil {
			return "", err
		}
	}
	return role, nil
}
