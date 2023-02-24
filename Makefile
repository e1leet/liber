CONFIG_PATH = ./configs/config.local.env
MIGRATIONS_PATH = migrations/
DATABASE_URI = postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable

.PHONY: run
run:
	go run ./cmd/app/main.go --config=${CONFIG_PATH}

.PHONY: lint
lint:
	golangci-lint run ./... --config=./.golangci.yml

.PHONY: test
test:
	go test -v -race -timeout=5m -cover ./...

.PHONY: test-coverage
test-coverage:
	go test -v -timeout=5m -covermode=count -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

.PHONY: swagger
swagger:
	swag init -g ./cmd/app/main.go

.PHONY: migrate.up
migrate.up:
	migrate -path $(MIGRATIONS_PATH) -database $(DATABASE_URI) up

.PHONY: migrate.down
migrate.down:
	migrate -path $(MIGRATIONS_PATH) -database $(DATABASE_URI) down