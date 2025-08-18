package config

import (
	"log"
	"os"
	"reflect"
	"sync"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	DBHost       string `env:"DB_HOST"     envRequired:"true"`
	DBName       string `env:"DB_NAME"     envRequired:"true"`
	DBPassword   string `env:"DB_PASSWORD" envRequired:"true"`
	DBPort       int    `env:"DB_PORT"     envRequired:"true"`
	DBUser       string `env:"DB_USER"     envRequired:"true"`
	DebugMode    bool   `env:"DEBUG_MODE"  envRequired:"false"`
	JWTSecretKey string `env:"JWT_SECRET_KEY"  envRequired:"true"`
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: .env file not found or failed to load, continuing with system env variables")
		}

		cfg := &Config{}
		if err := env.Parse(cfg); err != nil {
			log.Fatalf("Error loading environment variables: %v", err)
		}

		val := reflect.ValueOf(cfg).Elem()
		typ := reflect.TypeOf(cfg).Elem()

		for i := 0; i < val.NumField(); i++ {
			fieldType := typ.Field(i)
			requiredTag := fieldType.Tag.Get("envRequired")

			if requiredTag == "true" {
				envVar := os.Getenv(fieldType.Tag.Get("env"))
				if envVar == "" {
					log.Fatalf("Mandatory environment variable '%s' is not defined", fieldType.Tag.Get("env"))
				}
			}
		}

		instance = cfg
	})
	return instance
}
