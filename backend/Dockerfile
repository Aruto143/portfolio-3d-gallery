# 1. ビルド用コンテナ
FROM golang:1.24-alpine AS builder

WORKDIR /app

# 依存関係だけ先にコピー → キャッシュが効くようにする
COPY go.mod go.sum ./
RUN go mod download

# アプリ本体をコピーしてビルド
COPY . .
RUN go build -o server ./cmd/api

# 2. 実行用コンテナ（最小限）
FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/server .

# ポートは環境変数で管理してもOK
ENV PORT=8080

CMD ["./server"]
