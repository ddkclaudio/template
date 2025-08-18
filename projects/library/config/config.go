package config

import (
	"log"

	"github.com/ddkclaudio/envygo"
)

type Config struct {
	DBHost       string `env:"DB_HOST"       envRequired:"true"`
	DBName       string `env:"DB_NAME"       envRequired:"true"`
	DBPassword   string `env:"DB_PASSWORD"   envRequired:"true"`
	DBPort       int    `env:"DB_PORT"       envRequired:"true"`
	DBUser       string `env:"DB_USER"       envRequired:"true"`
	DebugMode    bool   `env:"DEBUG_MODE"    envRequired:"false"`
	JWTSecretKey string `env:"JWT_SECRET_KEY" envRequired:"true"`
}

var instance *Config

func NewConfig() *Config {
	if instance != nil {
		return instance
	}

	cfg, err := envygo.New[Config]()
	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}
	instance = cfg
	return instance
}

func GetConfig() *Config {

	if instance == nil {
		instance = NewConfig()
	}
	return instance
}
