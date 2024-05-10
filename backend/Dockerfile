# 開発用Dockerfile
FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


RUN go install github.com/cosmtrek/air@latest

COPY . .

# 秘密鍵と公開鍵をコピー
COPY key/private.pem /app/key/private.pem
COPY key/public.pem /app/key/public.pem

CMD ["air"]
