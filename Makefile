.PHONY: lint test-unit

lint:
	golangci-lint run ./...

lint-fix:
	golangci-lint run --fix ./...

test-unit:
	go test ./...

run:
	go run ./cmd/stats_server/