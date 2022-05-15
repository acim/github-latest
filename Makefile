.PHONY: lint test testv test-cov

lint:
	@golangci-lint run

test:
	@go test ./...

testv:
	@go test -v ./...

test-cov:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -func coverage.out
