NAME := knowledge-out
VERSION := v0.0.1

up :
	docker compose up

air:
	docker-compose exec app air -c .air.toml

gqlgen :
	cd backend/interface/gql && gqlgen generate

wire:
	go install github.com/google/wire/cmd/wire@latest
	cd app/adapter/http && wire

moq:
	go install github.com/matryer/moq@latest
	go generate ./app/application/...
	go generate ./app/domain/...

test:
	. ./.env && TEST_DB_USER=$${TEST_DB_USER} TEST_DB_PASSWORD=$${TEST_DB_PASSWORD} TEST_DB_HOST=$${TEST_DB_HOST} TEST_DB_PORT=$${TEST_DB_PORT} go test -v ./...