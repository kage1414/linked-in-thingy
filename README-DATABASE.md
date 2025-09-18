# Database Setup Guide

## Why Database Commands Are Not in package.json

Database operations should be separate from Node.js/npm scripts because:

1. **Separation of concerns**: Database is infrastructure, not frontend code
2. **Language independence**: Go applications shouldn't depend on npm for database setup
3. **Production deployment**: Production servers shouldn't need Node.js just to start the database
4. **Tool clarity**: Makes it clear this is a Go project with a React frontend

## Proper Database Management

### Using Makefile (Recommended)

```bash
# Start database
make db-up

# Stop database
make db-down

# Reset database (removes all data)
make db-reset

# Check database status
make db-status
```

### Using Shell Script

```bash
# Start database
./scripts/db.sh up

# Stop database
./scripts/db.sh down

# Reset database
./scripts/db.sh reset

# Check status
./scripts/db.sh status

# View logs
./scripts/db.sh logs
```

### Using Docker Compose Directly

```bash
# Start database
docker-compose up -d postgres

# Stop database
docker-compose down

# Reset database (removes all data)
docker-compose down -v && docker-compose up -d postgres
```

## Environment Variables

The application uses these environment variables for database configuration:

```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=jobboard
DB_SSLMODE=disable
```

## Complete Setup

```bash
# 1. Start database
make db-up

# 2. Install frontend dependencies
make frontend

# 3. Start development mode
make dev
```

This approach keeps database operations separate from frontend tooling and makes the project structure much cleaner.
