.PHONY: test

test:
	@echo "Running tests..."
	@go clean -testcache
	@go test ./...