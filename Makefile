.PHONY: test lint

test:
	go test ./... -coverprofile=coverage.out

lint:
	golangci-lint run -E gofumpt