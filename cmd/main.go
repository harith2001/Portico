package main

import (
	"log"

	"github.com/harith2001/portico/internal/config"
	"github.com/harith2001/portico/internal/gateway"
)

func main() {
	// Load configuration and connect to Redis
	config.ExampleClient_connect_basic()
	// Load the application configuration
	cfg := config.LoadConfig()

	app := gateway.NewApp(cfg)

	log.Printf("Starting Portico Gateway on %s...\n", cfg.Server.Port)
	if err := app.Listen(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
