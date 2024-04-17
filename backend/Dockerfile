# 開発用Dockerfile
FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/cosmtrek/air@latest

COPY . .

CMD ["air"]
