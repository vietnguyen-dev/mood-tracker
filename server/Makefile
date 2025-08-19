.PHONY: run dev build clean install-air

# Install air for hot reloading
install-air:
	go install github.com/cosmtrek/air@latest

# Run the server normally
run:
	go run main.go

# Run with hot reloading using air
dev: install-air
	air

# Build the binary
build:
	go build -o mood-tracker-server main.go

# Clean build artifacts
clean:
	rm -f mood-tracker-server
	go clean

# Install dependencies
deps:
	go mod tidy
	go mod download 