# Job Board Application

A modern job board application built with Go backend and React TypeScript frontend, featuring video streaming capabilities and hot reloading for development.

## Features

- **Backend**: Go with Gin web framework
- **Frontend**: React with TypeScript
- **API**: RESTful API with video streaming support
- **Video Streaming**: Built-in video streaming capabilities
- **Hot Reloading**: Both frontend and backend support hot reloading during development
- **Modern UI**: Beautiful, responsive design with modern UX practices

## Tech Stack

### Backend

- Go 1.21+
- Gin web framework
- CORS support
- Video streaming support

### Frontend

- React 18
- TypeScript
- Axios for API calls
- Modern CSS with responsive design

## Prerequisites

- Go 1.21 or higher
- Node.js 16 or higher
- npm or yarn
- Docker and Docker Compose (for PostgreSQL)

## Installation

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd linked-in-thingy
   ```

2. **Install Go dependencies**

   ```bash
   go mod tidy
   ```

3. **Install frontend dependencies**

   ```bash
   # Install frontend dependencies (note: you need to cd into frontend directory)
   cd frontend
   npm install
   cd ..
   ```

4. **Install development dependencies**

   ```bash
   npm install
   ```

5. **Install Air for Go hot reloading**
   ```bash
   cd ..
   go install github.com/air-verse/air@latest
   ```

## Development

### Start both frontend and backend with hot reloading

```bash
npm run dev
```

This will start:

- Go backend on `http://localhost:8080` with hot reloading
- React frontend on `http://localhost:3000` with hot reloading

### Individual services

**Backend only:**

```bash
npm run backend:dev
# or
air
```

**Frontend only:**

```bash
npm run frontend:start
```

## API Endpoints

### Jobs

- `GET /api/jobs` - Get all jobs
- `GET /api/jobs/:id` - Get job by ID
- `POST /api/jobs` - Create new job
- `PUT /api/jobs/:id` - Update job
- `DELETE /api/jobs/:id` - Delete job

### Videos

- `GET /api/videos` - Get all videos
- `GET /api/videos/:id` - Get video by ID
- `POST /api/videos` - Create new video

### Video Streaming

- `GET /video/:id` - Stream video by ID

## Project Structure

```
linked-in-thingy/
├── main.go                 # Main Go application
├── graph/                  # GraphQL schema and resolvers
│   ├── schema.graphqls     # GraphQL schema
│   ├── resolver.go         # GraphQL resolvers
│   └── model/              # Data models
├── frontend/               # React TypeScript frontend
│   ├── src/
│   │   ├── components/     # React components
│   │   ├── api.ts          # API service
│   │   ├── types.ts        # TypeScript types
│   │   └── App.tsx         # Main App component
│   └── public/             # Static assets
├── videos/                 # Video files directory
├── .air.toml              # Air configuration
└── package.json           # Node.js dependencies
```

## Building for Production

1. **Build the frontend**

   ```bash
   npm run frontend:build
   ```

2. **Build the Go application**

   ```bash
   npm run build
   ```

3. **Start the production server**
   ```bash
   npm start
   ```

The application will be available at `http://localhost:8080`

## Video Streaming

The application supports video streaming through the `/video/:id` endpoint. Place video files in the `videos/` directory with the format `{id}.mp4`.

## Development Notes

- The backend serves the React frontend in production
- CORS is configured to allow requests from `http://localhost:3000`
- Hot reloading is enabled for both frontend and backend during development
- The application uses mock data for demonstration purposes

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test your changes
5. Submit a pull request

## License

MIT License
