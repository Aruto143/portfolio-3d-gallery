package mysql

import (
	"context"

	usecase "github.com/Aruto143/portfolio-backend/internal/usecase/work"
	"github.com/jmoiron/sqlx"
)

// db上のカラム名に合わせたstruct
type workRow struct {
	ID           int64  `db:"id"`
	Title        string `db:"title"`
	Slug         string `db:"slug"`
	Summary      string `db:"summary"`
	ThumbnailURL string `db:"thumbnail_url"`
}

type WorkRepository struct {
	db *sqlx.DB
}

func NewWorkRepository(db *sqlx.DB) *WorkRepository {
	return &WorkRepository{db: db}
}

func (r *WorkRepository) List(ctx context.Context) ([]usecase.Work, error) {
	rows := []workRow{}
	err := r.db.SelectContext(ctx, &rows, `
		SELECT id, title, slug, summary, thumbnail_url
		  FROM works
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}

	result := make([]usecase.Work, 0, len(rows))
	for _, row := range rows {
		result = append(result, usecase.Work{
			ID:           row.ID,
			Title:        row.Title,
			Slug:         row.Slug,
			Summary:      row.Summary,
			ThumbnailURL: row.ThumbnailURL,
		})
	}
	return result, nil
}
