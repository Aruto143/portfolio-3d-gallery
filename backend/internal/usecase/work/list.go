package work

import "context"

// DomainのWorkを使うので、同じpackage名にしないように別名にすることも多い。
// ここでは簡略化して同じpackageに置いてしまうが、
// 実際には domain/work と usecase/work で package を分けてもOK。
type Work struct {
	ID          int64
	Title       string
	Slug        string
	Summary     string
	ThumbnailURL string
}

type WorkRepository interface {
	List(ctx context.Context) ([]Work, error)
}

// ListWorksUsecase は「作品を一覧取得する」ユースケース。
// この層は「どこからデータを取るか」（DBかAPIか）を知らない。
type ListWorksUsecase struct {
	repo WorkRepository
}

// NewListWorksUsecase は依存性（Repository）を注入するコンストラクタ。
func NewListWorksUsecase(repo WorkRepository) *ListWorksUsecase {
	return &ListWorksUsecase{repo: repo}
}

func (u *ListWorksUsecase) Run(ctx context.Context) ([]Work, error) {
	// 今は特にビジネスロジックはないが、
	// 将来「公開フラグのみ返す」「ソート条件を変える」
	// などをここに追加していける。
	return u.repo.List(ctx)
}
