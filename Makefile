# Job Board Application Makefile

.PHONY: help build run dev clean test db-up db-down db-reset setup

# Default target
help:
	@echo "Available commands:"
	@echo "  build      - Build the Go application"
	@echo "  run        - Run the application"
	@echo "  dev        - Run in development mode with hot reloading"
	@echo "  clean      - Clean build artifacts"
	@echo "  test       - Run tests"
	@echo "  db-up      - Start PostgreSQL database"
	@echo "  db-down    - Stop PostgreSQL database"
	@echo "  db-reset   - Reset database (remove data and restart)"
	@echo "  setup      - Complete setup (database + frontend)"
	@echo "  frontend   - Install frontend dependencies"

# Build the application
build:
	@echo "Building job board application..."
	go build -o job-board main.go
	@echo "Build complete: ./job-board"

# Run the application
run: build
	@echo "Starting job board application..."
	./job-board

# Development mode with hot reloading
dev:
	@echo "Starting development mode..."
	@echo "Backend will run on :8080 with hot reloading"
	@echo "Frontend will run on :3000 with hot reloading"
	@echo "Make sure PostgreSQL is running (make db-up)"
	npm run dev

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f job-board
	rm -rf frontend/build
	@echo "Clean complete"

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Database commands
db-up:
	@echo "Starting PostgreSQL database..."
	docker-compose up -d postgres
	@echo "Database started. Waiting for it to be ready..."
	@sleep 5
	@echo "Database is ready!"

db-down:
	@echo "Stopping PostgreSQL database..."
	docker-compose down

db-reset:
	@echo "Resetting database (removing all data)..."
	docker-compose down -v
	docker-compose up -d postgres
	@echo "Database reset complete. Waiting for it to be ready..."
	@sleep 5
	@echo "Database is ready!"

# Frontend setup
frontend:
	@echo "Installing frontend dependencies..."
	cd frontend && npm install
	@echo "Frontend dependencies installed"

# Complete setup
setup: db-up frontend
	@echo "Setup complete!"
	@echo "Run 'make dev' to start development mode"
	@echo "Or run 'make run' to start the application"
