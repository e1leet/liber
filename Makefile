CONFIG_PATH = ./configs/config.local.env

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