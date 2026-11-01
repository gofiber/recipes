package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/spf13/viper"
)

// Config holds the application configuration loaded from defaults, files, and environment variables.
type Config struct {
	Port int    `mapstructure:"port"`
	Env  string `mapstructure:"env"`
	App  struct {
		Name    string `mapstructure:"name"`
		Version string `mapstructure:"version"`
	} `mapstructure:"app"`
}

// loadConfig initializes Viper, configures default values, reads config.yaml if present,
// maps environment variables prefixed with APP_, and unmarshals the result into a Config struct.
func loadConfig() (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	// Set defaults
	v.SetDefault("port", 3000)
	v.SetDefault("env", "development")
	v.SetDefault("app.name", "Fiber Viper Demo")
	v.SetDefault("app.version", "1.0.0")

	// Read environment variables (e.g. APP_PORT=8080 or APP_APP_NAME="Custom Name")
	v.SetEnvPrefix("APP")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// Read config file if present
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
		log.Println("Config file not found, using default and environment variables")
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &cfg, nil
}

// main boots the Fiber web server using the loaded configuration.
func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"app":     cfg.App.Name,
			"version": cfg.App.Version,
			"env":     cfg.Env,
			"port":    cfg.Port,
		})
	})

	app.Get("/healthz", func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Starting %s on %s in %s mode\n", cfg.App.Name, addr, cfg.Env)
	log.Fatal(app.Listen(addr))
}
