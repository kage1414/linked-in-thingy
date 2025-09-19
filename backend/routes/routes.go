package routes

import (
	"job-board/backend/config"
	"job-board/backend/handlers"
	"job-board/backend/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all routes for the application
func SetupRoutes(h *handlers.Handler, cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// Add middleware
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.RecoveryMiddleware())
	r.Use(middleware.RequestIDMiddleware())
	r.Use(middleware.SecurityMiddleware())
	r.Use(middleware.RateLimitMiddleware())
	r.Use(middleware.CORSMiddleware())

	// API routes
	api := r.Group("/api")
	{
		// Job routes
		api.GET("/jobs", h.GetJobs)
		api.GET("/jobs/:id", h.GetJob)
		api.POST("/jobs", h.CreateJob)
		api.PUT("/jobs/:id", h.UpdateJob)
		api.DELETE("/jobs/:id", h.DeleteJob)

		// Video routes
		api.GET("/videos", h.GetVideos)
		api.GET("/videos/:id", h.GetVideo)
		api.POST("/videos", h.CreateVideo)
	}

	// Video streaming route
	r.GET("/video/:id", h.StreamVideo)

	// Static file serving for React frontend
	r.Static("/static", "./frontend/build/static")
	r.StaticFile("/", "./frontend/build/index.html")
	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/build/index.html")
	})

	return r
}
