.PHONY: lint check test testv test-cov update

lint:
	@golangci-lint run

# Mirrors the static gates of ectobit/reusable-workflows go-check.yaml in the
# same order; update both together. Run before every push, with the affected
# tests. Tests are separate because CI runs them in its own job.
check: lint
	govulncheck ./...
	go fix -diff ./...

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
