.PHONY: lint test testv test-cov update

lint:
	@golangci-lint run

test:
	@go test ./...

testv:
	@go test -v ./...

test-cov:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -func coverage.out

update:
	@go get -u
	@go mod tidy
