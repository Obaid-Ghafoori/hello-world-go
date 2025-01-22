.PHONY: all test coverage lint vet fmt deps init

# Default target
all: init deps test coverage lint vet fmt

# Initialize Go module
init:
	go mod init github.com/Obaid-Ghafoori/hello-world-go || true
	go mod tidy

# Install dependencies
deps:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install golang.org/x/lint/golint@latest
	go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest

# Run tests with coverage
test:
	go test -v -race -coverprofile=coverage.out ./...

coverage:
	go tool cover -html=coverage.out -o coverage.html

# Run linting
lint:
	$(shell go env GOPATH)/bin/staticcheck ./...

# Run go vet
vet:
	go vet ./...

# Format code
fmt:
	if [ -n "$$(gofmt -l .)" ]; then \
		echo "These files are not formatted:"; \
		gofmt -l .; \
		echo "Please run 'go fmt ./...' to format your code"; \
		exit 1; \
	fi

# Security checks
security:
	gosec ./...

# Clean build artifacts
clean:
	go clean
	rm -f coverage.out coverage.html
	rm -f go.mod go.sum 