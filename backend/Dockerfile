FROM golang:1.22 AS builder

WORKDIR /app

# go.mod と go.sum をコピー
COPY go.mod go.sum ./
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# スタティックバイナリをビルド
ENV CGO_ENABLED=0
RUN go build -o main ./cmd/app/main.go

# 実行ステージ
FROM alpine:3.15

WORKDIR /app

# ビルドステージから実行ファイルをコピー
COPY --from=builder /app/main .

CMD ["/app/main"]