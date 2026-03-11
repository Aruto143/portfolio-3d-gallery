# portfolio-3d-gallery

作品一覧・作品詳細を表示する Web アプリケーションです。

<!-- 見た目や機能を作るだけではなく、**構成・設計・開発基盤をどのように考えて組み立てたか**を重視して実装しています。 -->

<!-- ## このプロジェクトで重視していること

- フロントエンド / バックエンド / DB の責務分離
- Go バックエンドにおけるレイヤー分離
- PostgreSQL を用いた永続化
- Docker Compose によるローカル開発環境の再現性
- Makefile による開発操作の統一
- 将来的な CI/CD・デプロイを見据えた構成

--- -->

## 現在できていること

- 作品一覧 API
  - `GET /api/works`
- 作品詳細 API
  - `GET /api/works/:slug`
- Next.js による作品一覧ページ
  - `/works`
- Next.js による作品詳細ページ
  - `/works/[slug]`
- Docker Compose による `frontend / backend / db` の統合起動
- Makefile による開発操作の統一
- GitHub Actions による最小 CI
  - backend: `go vet`, `go test ./...`
  - frontend: `npm run build`
- デプロイ構成
  - frontend: Vercel
  - backend: Render
  - database: Render Postgres

---

## 技術スタック

### Frontend

- Next.js
- React
- TypeScript
- Tailwind CSS

### Backend

- Go
- Gin
- sqlx
- pgx

### Database

- PostgreSQL

### Development Environment

- Docker
- Docker Compose
- Makefile

---

## アーキテクチャ概要

このプロジェクトでは、役割ごとに責務を分けることを意識しています。

### 全体構成

```text
Browser
  ↓
frontend (Next.js)
  ↓
backend (Go API)
  ↓
db (PostgreSQL)
```

### Docker Compose 上での構成

```text
Browser
  ↓ localhost:3000
frontend container
  ↓ http://backend:18080
backend container
  ↓ postgres://appuser:apppass@db:5432/portfolio
db container
```

### バックエンドのレイヤー構成

```text
backend/
  cmd/api/main.go                    # アプリケーションの組み立てと起動
  internal/
    usecase/work/                    # ユースケース
    interface/http/                  # HTTP ハンドラ
    infrastructure/postgres/         # PostgreSQL Repository 実装
    infra/db/                        # DB 接続
```

#### 各レイヤーの役割

```text
- `cmd/api/main.go`
  - DB 接続、Repository、Usecase、Handler を組み立てて起動する
- `usecase`
  - アプリケーションが何をしたいかを表現する
- `interface/http`
  - HTTPリクエスト/レスポンスを扱う
- `infrastructure/postgres`
  - PostgreSQLに対する具体的なデータアクセスを実装する
- `infra/db`
  - DB接続設定をまとめる
```

---

## 今後やりたいこと

- README への構成図・設計意図の追記
- トップページの実装
- 管理画面の追加
- 認証機能の追加
