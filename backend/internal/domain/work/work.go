package work

import "time"

// Work は「ギャラリーに並べる1つの作品」を表すドメインモデル
// ここではDBやフレームワークの事情を一切考えない。
type Work struct {
	ID          int64
	Title       string
	Slug        string
	Summary     string
	Description string
	ThumbnailURL string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
