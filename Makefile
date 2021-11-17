.PHONY: lint test test-verbose test-with-coverage

lint:
	@golangci-lint run --exclude-use-default=false --enable-all \
		--disable golint \
		--disable interfacer \
		--disable scopelint \
		--disable maligned

test:
	@go test ./...

test-verbose:
	@go test -v ./...

test-with-coverage:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -func coverage.out
