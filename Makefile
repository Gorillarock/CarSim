
## Run Command: make test

# Tells make that "test" is not a file
.PHONY: test

test:
# Clean the test cache
	go clean -testcache
# Run all tests in the current directory and all subdirectories
	go test ./...