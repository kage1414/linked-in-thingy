package server

import (
	"log"
	"os"

	"job-board/backend/config"
	"job-board/backend/database"
	"job-board/backend/handlers"
	"job-board/backend/routes"
	"job-board/backend/streaming"
)

// Server represents the application server
type Server struct {
	config *config.Config
}

// NewServer creates a new server instance
func NewServer() *Server {
	return &Server{
		config: config.LoadConfig(),
	}
}

// Start initializes and starts the server
func (s *Server) Start() error {
	// Initialize database
	if err := database.ConnectDatabase(s.config.Database.URL); err != nil {
		return err
	}

	// Run migrations
	if err := database.MigrateDatabase(); err != nil {
		return err
	}

	// Seed database with sample data
	if err := database.SeedDatabase(); err != nil {
		return err
	}

	// Initialize services
	jobService := database.NewJobService(database.DB)
	videoService := database.NewVideoService(database.DB)
	videoStreamer := streaming.NewVideoStreamer(s.config.Video.Directory)

	// Initialize handlers
	handler := handlers.NewHandler(jobService, videoService, videoStreamer)

	// Setup routes
	router := routes.SetupRoutes(handler, s.config)

	// Create videos directory if it doesn't exist
	if err := os.MkdirAll(s.config.Video.Directory, 0o755); err != nil {
		log.Printf("Warning: Failed to create video directory: %v", err)
	}

	// Start server
	addr := s.config.Server.Host + ":" + s.config.Server.Port
	log.Printf("Server starting on %s", addr)
	log.Printf("API available at: http://%s/api", addr)

	return router.Run(addr)
}
