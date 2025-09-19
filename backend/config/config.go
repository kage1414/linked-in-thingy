package config

import (
	"os"
	"strconv"
)

// Config holds all configuration for our application
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	CORS     CORSConfig
	Video    VideoConfig
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port string
	Host string
}

// DatabaseConfig holds database-related configuration
type DatabaseConfig struct {
	URL string
}

// CORSConfig holds CORS-related configuration
type CORSConfig struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
}

// VideoConfig holds video-related configuration
type VideoConfig struct {
	Directory string
}

// LoadConfig loads configuration from environment variables with defaults
func LoadConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
			Host: getEnv("HOST", "localhost"),
		},
		Database: DatabaseConfig{
			URL: getEnv("DATABASE_URL", "host=localhost user=postgres password=postgres dbname=jobboard port=5432 sslmode=disable TimeZone=America/New_York"),
		},
		CORS: CORSConfig{
			AllowOrigins:     getEnvSlice("CORS_ALLOW_ORIGINS", []string{"http://localhost:3000"}),
			AllowMethods:     getEnvSlice("CORS_ALLOW_METHODS", []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			AllowHeaders:     getEnvSlice("CORS_ALLOW_HEADERS", []string{"Origin", "Content-Type", "Accept", "Authorization"}),
			ExposeHeaders:    getEnvSlice("CORS_EXPOSE_HEADERS", []string{"Content-Length"}),
			AllowCredentials: getEnvBool("CORS_ALLOW_CREDENTIALS", true),
		},
		Video: VideoConfig{
			Directory: getEnv("VIDEO_DIRECTORY", "videos"),
		},
	}
}

// getEnv gets an environment variable with a fallback default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvSlice gets an environment variable as a slice with a fallback default value
func getEnvSlice(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		// Simple comma-separated parsing - in production you might want more sophisticated parsing
		return []string{value}
	}
	return defaultValue
}

// getEnvBool gets an environment variable as a boolean with a fallback default value
func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.ParseBool(value); err == nil {
			return parsed
		}
	}
	return defaultValue
}
