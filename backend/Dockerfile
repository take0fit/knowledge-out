FROM golang:1.22 AS builder

WORKDIR /app

# go.mod と go.sum をコピー
COPY go.mod go.sum ./
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# スタティックバイナリをビルド（クロスコンパイル対応）
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o main ./cmd/api-server/main.go
RUN go build -o migration ./cmd/migration/migration.go

RUN chmod +x main migration

# 実行ステージ
FROM alpine:3.15

WORKDIR /app

# ビルドステージから実行ファイルをコピー
COPY --from=builder /app/main .
COPY --from=builder /app/migration .

CMD ["/app/main"]
