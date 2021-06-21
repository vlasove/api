.PHONY: build, lint, test, migrateup, migratedown, doc

build:
	go build -v ./cmd/apiserver

lint:
	golint ./... && golangci-lint run

test:
	go test -v -race ./...

migrateup:
	migrate -path migrations -database "postgres://admin:admin@localhost:5432/restapi_dev?sslmode=disable" up

migratedown:
	migrate -path migrations -database "postgres://admin:admin@localhost:5432/restapi_dev?sslmode=disable" down

migrateup_test:
	migrate -path migrations -database "postgres://admin:admin@localhost:5432/restapi_test?sslmode=disable" up

migratedown_test:
	migrate -path migrations -database "postgres://admin:admin@localhost:5432/restapi_test?sslmode=disable" down

doc:
	godoc -http :8000

.DEFAULT_GOAL := build