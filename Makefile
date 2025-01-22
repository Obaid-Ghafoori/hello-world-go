.PHONY: test lint vet fmt security clean

# Default target
all: test lint vet fmt security

# Download dependencies
deps:
	go mod download
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install golang.org/x/lint/golint@latest
	go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest

# Run tests with coverage
test:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Run linting
lint:
	golint -set_exit_status ./...
	staticcheck ./...

# Run go vet
vet:
	go vet ./...

# Format code
fmt:
	goimports -w .

# Security checks
security:
	gosec ./...

# Clean build artifacts
clean:
	go clean
	rm -f coverage.out coverage.html 