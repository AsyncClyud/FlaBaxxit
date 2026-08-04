package config

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	DBHost            string
	DBPort            string
	DBUser            string
	DBPassword        string
	DBName            string
	DBSSLMode         string
	JWTSecret         string
	Cloudflare_secret string
	Port              string
}

// Loading config from .env file.
func Load() *Config {

	cfg := &Config{
		DBHost:            getEnv("DB_HOST", "localhost"),
		DBPort:            getEnv("DB_PORT", "5432"),
		DBUser:            os.Getenv("DB_USER"),
		DBPassword:        os.Getenv("DB_PASSWORD"),
		DBName:            os.Getenv("DB_NAME"),
		DBSSLMode:         getEnv("DB_SSLMODE", "disable"),
		JWTSecret:         os.Getenv("JWT_SECRET"),
		Cloudflare_secret: os.Getenv("CLOUDFLARE_SECRET"),
		Port:              getEnv("PORT", "8080"),
	}

	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET is required")
	}

	return cfg
}

// Returns database connection string
func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode,
	)
}

// Returns redis URL
func (c *Config) RedisURL() string {
	return os.Getenv("REDIS_URL")
}

// Get env by argument
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
