package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// GetDatabaseConfig returns database configuration from environment variables
func GetDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "jobboard"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}

// ConnectDatabase establishes connection to PostgreSQL database
func ConnectDatabase() error {
	config := GetDatabaseConfig()
	
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL database")
	return nil
}

// MigrateDatabase runs database migrations
func MigrateDatabase() error {
	if DB == nil {
		return fmt.Errorf("database connection not established")
	}

	// Auto-migrate the schema
	err := DB.AutoMigrate(&Job{}, &Video{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database migration completed successfully")
	return nil
}

// SeedDatabase populates the database with initial data
func SeedDatabase() error {
	if DB == nil {
		return fmt.Errorf("database connection not established")
	}

	// Check if jobs already exist
	var count int64
	DB.Model(&Job{}).Count(&count)
	if count > 0 {
		log.Println("Database already seeded, skipping...")
		return nil
	}

	// Create sample jobs
	jobs := []Job{
		{
			Title:       "Senior Software Engineer",
			Company:     "Tech Corp",
			Description: "We are looking for a senior software engineer to join our team. You will be responsible for designing and implementing scalable web applications using modern technologies.",
			Location:    "San Francisco, CA",
			Salary:      stringPtr("$120,000 - $150,000"),
			Requirements: []string{
				"5+ years of experience in software development",
				"Proficiency in Go and React",
				"Experience with GraphQL and REST APIs",
				"Knowledge of PostgreSQL and database design",
				"Experience with Docker and Kubernetes",
			},
			Benefits: []string{
				"Health insurance",
				"401k matching up to 6%",
				"Flexible work hours",
				"Remote work options",
				"Professional development budget",
			},
			VideoURL: stringPtr("/video/1"),
		},
		{
			Title:       "Frontend Developer",
			Company:     "StartupXYZ",
			Description: "Join our fast-growing startup as a frontend developer. You'll work on building beautiful, responsive user interfaces and contribute to our product development.",
			Location:    "Remote",
			Salary:      stringPtr("$80,000 - $100,000"),
			Requirements: []string{
				"3+ years of React experience",
				"TypeScript proficiency",
				"Experience with modern CSS frameworks",
				"Knowledge of state management (Redux, Zustand)",
				"Experience with testing frameworks (Jest, Cypress)",
			},
			Benefits: []string{
				"Stock options",
				"Unlimited PTO",
				"Learning budget of $2,000/year",
				"Home office stipend",
				"Flexible schedule",
			},
			VideoURL: stringPtr("/video/2"),
		},
		{
			Title:       "DevOps Engineer",
			Company:     "CloudTech Solutions",
			Description: "We're seeking a DevOps engineer to help us scale our infrastructure and improve our deployment processes. You'll work with AWS, Kubernetes, and modern CI/CD tools.",
			Location:    "New York, NY",
			Salary:      stringPtr("$110,000 - $140,000"),
			Requirements: []string{
				"4+ years of DevOps experience",
				"Strong AWS knowledge",
				"Experience with Kubernetes",
				"Knowledge of Terraform or CloudFormation",
				"Experience with CI/CD pipelines",
			},
			Benefits: []string{
				"Comprehensive health coverage",
				"401k with company matching",
				"Annual bonus potential",
				"Conference attendance budget",
				"Gym membership",
			},
		},
	}

	// Create jobs in database
	for _, job := range jobs {
		if err := DB.Create(&job).Error; err != nil {
			return fmt.Errorf("failed to create job: %w", err)
		}
	}

	// Create sample videos
	videos := []Video{
		{
			JobID:     1,
			Title:     "Company Culture Video",
			URL:       "/video/1",
			Duration:  intPtr(120),
			Thumbnail: stringPtr("/thumbnails/1.jpg"),
		},
		{
			JobID:     2,
			Title:     "Team Introduction",
			URL:       "/video/2",
			Duration:  intPtr(90),
			Thumbnail: stringPtr("/thumbnails/2.jpg"),
		},
	}

	// Create videos in database
	for _, video := range videos {
		if err := DB.Create(&video).Error; err != nil {
			return fmt.Errorf("failed to create video: %w", err)
		}
	}

	log.Println("Database seeded with sample data successfully")
	return nil
}

// Helper functions
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}
