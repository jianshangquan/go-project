# Makefile

# Define your binary name
BINARY_NAME=myapp

# Set Go build flags (optional)
GOFLAGS := -v

# Command for running the Go app in development mode
dev:
	@echo "Running Go development server..."
	go run src/main.go

# Command for building the Go binary
build:
	@echo "Building the Go project..."
	go build $(GOFLAGS) -o $(BINARY_NAME) src/main.go

# Command for cleaning build artifacts (e.g., binary)
clean:
	@echo "Cleaning build artifacts..."
	rm -f $(BINARY_NAME)

# Command for running tests
test:
	@echo "Running tests..."
	go test -v ./...

# Command for formatting Go code
fmt:
	@echo "Formatting Go code..."
	go fmt ./...

# Command for installing dependencies (if you have any)
deps:
	@echo "Installing dependencies..."
	go mod tidy

# Command for running a development environment (with hot reloading using air)
dev-reload:
	@echo "Running development environment with air..."
	air
