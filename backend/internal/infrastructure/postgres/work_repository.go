package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"

	usecase "github.com/Aruto143/portfolio-backend/internal/usecase/work"
)

type WorkRepository struct {
	db *sqlx.DB
}

type workRow struct {
	ID           int64  `db:"id"`
	Title        string `db:"title"`
	Slug         string `db:"slug"`
	Summary      string `db:"summary"`
	ThumbnailURL string `db:"thumbnail_url"`
}

func NewWorkRepository(db *sqlx.DB) usecase.WorkRepository {
	return &WorkRepository{db: db}
}

func (r *WorkRepository) List(ctx context.Context) ([]usecase.Work, error) {
	rows := []workRow{}

	query := `
		SELECT
			id,
			title,
			slug,
			summary,
			thumbnail_url
		FROM works
		ORDER BY id DESC
	`

	if err := r.db.SelectContext(ctx, &rows, query); err != nil {
		return nil, err
	}

	works := make([]usecase.Work, 0, len(rows))
	for _, row := range rows {
		works = append(works, usecase.Work{
			ID:           row.ID,
			Title:        row.Title,
			Slug:         row.Slug,
			Summary:      row.Summary,
			ThumbnailURL: row.ThumbnailURL,
		})
	}

	return works, nil
}