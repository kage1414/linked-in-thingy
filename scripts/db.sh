#!/bin/bash

# Database management script for Job Board application

set -e

case "$1" in
    up)
        echo "Starting PostgreSQL database..."
        docker-compose up -d postgres
        echo "Database started. Waiting for it to be ready..."
        sleep 5
        echo "Database is ready!"
        ;;
    down)
        echo "Stopping PostgreSQL database..."
        docker-compose down
        ;;
    reset)
        echo "Resetting database (removing all data)..."
        docker-compose down -v
        docker-compose up -d postgres
        echo "Database reset complete. Waiting for it to be ready..."
        sleep 5
        echo "Database is ready!"
        ;;
    status)
        echo "Checking database status..."
        docker-compose ps postgres
        ;;
    logs)
        echo "Showing database logs..."
        docker-compose logs postgres
        ;;
    *)
        echo "Usage: $0 {up|down|reset|status|logs}"
        echo ""
        echo "Commands:"
        echo "  up     - Start PostgreSQL database"
        echo "  down   - Stop PostgreSQL database"
        echo "  reset  - Reset database (remove data and restart)"
        echo "  status - Show database status"
        echo "  logs   - Show database logs"
        exit 1
        ;;
esac
