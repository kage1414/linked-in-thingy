package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"job-board/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	if err := database.ConnectDatabase(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run migrations
	if err := database.MigrateDatabase(); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Seed database with sample data
	if err := database.SeedDatabase(); err != nil {
		log.Fatal("Failed to seed database:", err)
	}

	// Initialize services
	jobService := database.NewJobService(database.DB)
	videoService := database.NewVideoService(database.DB)

	// Setup Gin router
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// API routes
	api := r.Group("/api")
	{
		api.GET("/jobs", func(c *gin.Context) { getJobs(c, jobService) })
		api.GET("/jobs/:id", func(c *gin.Context) { getJob(c, jobService) })
		api.POST("/jobs", func(c *gin.Context) { createJob(c, jobService) })
		api.PUT("/jobs/:id", func(c *gin.Context) { updateJob(c, jobService) })
		api.DELETE("/jobs/:id", func(c *gin.Context) { deleteJob(c, jobService) })
		api.GET("/videos", func(c *gin.Context) { getVideos(c, videoService) })
		api.GET("/videos/:id", func(c *gin.Context) { getVideo(c, videoService) })
		api.POST("/videos", func(c *gin.Context) { createVideo(c, videoService) })
	}

	// Video streaming endpoint
	r.GET("/video/:id", func(c *gin.Context) {
		videoID := c.Param("id")
		// In a real app, you'd fetch video info from database
		// For now, we'll serve a sample video
		videoPath := filepath.Join("videos", videoID+".mp4")

		if _, err := os.Stat(videoPath); os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
			return
		}

		c.Header("Content-Type", "video/mp4")
		c.Header("Accept-Ranges", "bytes")
		c.File(videoPath)
	})

	// Serve React frontend
	r.Static("/static", "./frontend/build/static")
	r.StaticFile("/", "./frontend/build/index.html")
	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/build/index.html")
	})

	// Create videos directory if it doesn't exist
	os.MkdirAll("videos", 0o755)

	log.Println("Server starting on :8080")
	log.Println("API available at: http://localhost:8080/api")
	log.Fatal(r.Run(":8080"))
}

// API handlers
func getJobs(c *gin.Context, jobService *database.JobService) {
	jobs, err := jobService.GetAllJobs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch jobs"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": jobs})
}

func getJob(c *gin.Context, jobService *database.JobService) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	job, err := jobService.GetJobByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": job})
}

func createJob(c *gin.Context, jobService *database.JobService) {
	var job database.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := jobService.CreateJob(&job); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create job"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": job})
}

func updateJob(c *gin.Context, jobService *database.JobService) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	var job database.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := jobService.UpdateJob(uint(id), &job); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": job})
}

func deleteJob(c *gin.Context, jobService *database.JobService) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	if err := jobService.DeleteJob(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func getVideos(c *gin.Context, videoService *database.VideoService) {
	videos, err := videoService.GetAllVideos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch videos"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": videos})
}

func getVideo(c *gin.Context, videoService *database.VideoService) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid video ID"})
		return
	}

	video, err := videoService.GetVideoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": video})
}

func createVideo(c *gin.Context, videoService *database.VideoService) {
	var video database.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := videoService.CreateVideo(&video); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create video"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": video})
}
