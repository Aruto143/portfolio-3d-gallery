package work

import "context"

type FindWorkBySlugUsecase struct {
	repo WorkRepository
}

func NewFindWorkBySlugUsecase(repo WorkRepository) *FindWorkBySlugUsecase {
	return &FindWorkBySlugUsecase{repo: repo}
}

func (u *FindWorkBySlugUsecase) Run(ctx context.Context, slug string) (*Work, error) {
	return u.repo.FindBySlug(ctx, slug)
}