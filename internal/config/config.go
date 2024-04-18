package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type EnvType string

const (
	EnvLocal EnvType = "local"
	EnvProd  EnvType = "prod"
)

type (
	Config struct {
		Env        EnvType `yaml:"env" env-required:"true"`
		DB         `yaml:"db"`
		Cache      `yaml:"cache"`
		HTTPServer `yaml:"http_server"`
	}
	DB struct {
		Host         string `env:"POSTGRES_HOST" env-default:"auth-postgres"`
		Name         string `env:"POSTGRES_NAME" env-default:"auth"`
		User         string `env:"POSTGRES_USER"`
		Pass         string `env:"POSTGRES_PASSWORD"`
		MaxOpenConns int    `yaml:"max_open_conns" env-default:"5"`
		MaxIdleConns int    `yaml:"max_idle_conns" env-default:"10"`
	}
	HTTPServer struct {
		RunPort     string        `ENV:"RUN_PORT" env-default:"8080"`
		Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
		IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
		LogFileName string        `yaml:"log_file_name" env-default:"service.log"`
	}
	Cache struct {
		Host string `yaml:"host" env-default:"auth-redis"`
		Pass string `env:"REDIS_PASSWORD"`
	}
)

func NewConfig() *Config {
	cfg := Config{}
	err := cleanenv.ReadConfig("./config/config.yaml", &cfg)
	if err != nil {
		log.Fatalf("could not read config file: %v", err)
	}
	if cfg.Env != EnvLocal && cfg.Env != EnvProd {
		log.Fatalf("unknown env: %s", cfg.Env)
	}

	err = godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("could not load .env file: %s\n", err)
	}
	cfg.DB.User = readEnvVar("POSTGRES_USER")
	cfg.DB.Pass = readEnvVar("POSTGRES_PASSWORD")
	cfg.Cache.Pass = readEnvVar("REDIS_PASSWORD")
	return &cfg
}

func readEnvVar(name string) (value string) {
	value = os.Getenv(name)
	if value == "" {
		log.Fatalf("env variable %s not found\n", name)
	}
	return
}
