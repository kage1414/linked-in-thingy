# Development Guide

## Hot Reloading Setup

### Frontend Hot Reloading

The frontend uses React's built-in hot reloading with Create React App. It runs on port 3000 and proxies API requests to the backend on port 8080.

### Backend Hot Reloading

The backend uses Air for hot reloading. It watches for Go file changes and automatically rebuilds and restarts the server.

## Development Commands

### Start Everything (Recommended)

```bash
# Start database
make db-up

# Start both frontend and backend with hot reloading
npm run dev
```

This will start:

- Backend on `http://localhost:8080` (with hot reloading)
- Frontend on `http://localhost:3000` (with hot reloading)
- Frontend will proxy API requests to backend

### Individual Services

**Backend only:**

```bash
make run          # Run once
make backend:dev  # Run with hot reloading
```

**Frontend only:**

```bash
make frontend:start
```

## Troubleshooting Hot Reloading

### Frontend Not Reloading

1. Check if the frontend is running on port 3000
2. Make sure no other process is using port 3000
3. Try restarting: `Ctrl+C` then `npm run dev`

### Backend Not Reloading

1. Check if Air is installed: `which air`
2. Make sure the backend is running on port 8080
3. Check Air logs for errors

### API Connection Issues

1. Make sure backend is running on port 8080
2. Check browser network tab for failed requests
3. Verify proxy configuration in `frontend/package.json`

## File Structure

```
job-board/
├── main.go              # Backend server
├── database/            # Database models and services
├── frontend/            # React frontend
│   ├── src/
│   │   ├── components/  # React components
│   │   ├── api.ts       # API client
│   │   └── types.ts     # TypeScript types
│   └── package.json     # Frontend dependencies + proxy config
├── .air.toml           # Air configuration
└── package.json        # Root package.json with dev scripts
```

## Environment Variables

### Backend

```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=jobboard
DB_SSLMODE=disable
```

### Frontend

```bash
REACT_APP_API_URL=http://localhost:8080/api  # Only needed in production
```

## Common Issues

1. **Port conflicts**: Make sure ports 3000 and 8080 are free
2. **Database connection**: Ensure PostgreSQL is running (`make db-up`)
3. **CORS errors**: Backend has CORS configured for localhost:3000
4. **Hot reload not working**: Restart the dev server
