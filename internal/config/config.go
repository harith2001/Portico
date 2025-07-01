package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Port string
	}
	JWTSecret string
	MongoURI  string
	RedisAddr string
}

func LoadConfig() Config {
	_ = godotenv.Load() // Load .env into os.Environ

	cfg := Config{}
	cfg.Server.Port = getEnv("PORT", "8080")
	cfg.JWTSecret = getEnv("JWT_SECRET", "hardToGetGuessTHISPASSWORD")
	cfg.MongoURI = getEnv("MONGO_URI", "mongodb://localhost:27017")
	cfg.RedisAddr = getEnv("REDIS_ADDR", "localhost:6379")

	return cfg
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Printf("Missing env var %s, defaulting to %s", key, fallback)
		return fallback
	}
	return val
}
