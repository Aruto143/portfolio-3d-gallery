package work

import (
	"context"
	"testing"
)

type mockWorkRepository struct {
	listResult []Work
	listErr    error
}

func (m *mockWorkRepository) List(ctx context.Context) ([]Work, error) {
	return m.listResult, m.listErr
}

func (m *mockWorkRepository) FindBySlug(ctx context.Context, slug string) (*Work, error) {
	for _, w := range m.listResult {
		if w.Slug == slug {
			work := w
			return &work, nil
		}
	}
	return nil, nil
}

func TestListWorksUsecase_Run(t *testing.T) {
	repo := &mockWorkRepository{
		listResult: []Work{
			{
				ID:           1,
				Title:        "3D Portfolio Gallery",
				Slug:         "3d-portfolio-gallery",
				Summary:      "Three.js を使った3Dギャラリー",
				ThumbnailURL: "https://example.com/thumb1.png",
			},
		},
	}

	usecase := NewListWorksUsecase(repo)

	works, err := usecase.Run(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(works) != 1 {
		t.Fatalf("expected 1 work, got %d", len(works))
	}

	if works[0].Slug != "3d-portfolio-gallery" {
		t.Fatalf("unexpected slug: %s", works[0].Slug)
	}
}

func TestFindWorkBySlugUsecase_Run(t *testing.T) {
	repo := &mockWorkRepository{
		listResult: []Work{
			{
				ID:           2,
				Title:        "Task Management API",
				Slug:         "task-management-api",
				Summary:      "Clean Architecture を意識したAPI設計のサンプル",
				ThumbnailURL: "https://example.com/thumb2.png",
			},
		},
	}

	usecase := NewFindWorkBySlugUsecase(repo)

	work, err := usecase.Run(context.Background(), "task-management-api")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if work == nil {
		t.Fatal("expected work, got nil")
	}

	if work.Title != "Task Management API" {
		t.Fatalf("unexpected title: %s", work.Title)
	}
}
