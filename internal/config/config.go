// Package config config
package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config struct of config APP.
type Config struct {
	Postgres `yaml:"postgres"`
	App      `yaml:"app"`
	Logger   `yaml:"logger"`
}

// Postgres Db config.
type Postgres struct {
	Host     string `yaml:"host" env:"PSQL_HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"PSQL_PORT" env-default:"5432"`
	Username string `yaml:"username" env:"PSQL_USERNAME" env-default:"pg"`
	Password string `yaml:"password" env:"PSQL_PASSWORD" env-default:"pass"`
	DBName   string `yaml:"dbName" env:"PSQL_DB_NAME" env-default:"crud"`
	SSLMode  string `yaml:"sslMode" env:"PSQL_SSL_MODE" env-default:"disable"`
}

// App config.
type App struct {
	Port           string `yaml:"port" env:"APP_PORT" env-default:"8000"`
	MaxHeaderBytes int    `yaml:"maxHeaderBytes" env:"MAX_HEADER_BYTES" env-default:"20"`
	WriteTimeout   int    `yaml:"writeTimeout" env:"APP_WRITE_TIMEOUT" env-default:"10"`
	ReadTimeout    int    `yaml:"readTimeout" env:"APP_READ_TIMEOUT" env-default:"10"`
}

// Logger config.
type Logger struct {
	LogLevel int `yaml:"logLevel" env:"LOG_LEVEL" env-default:"0"`
}

// GetConfig parse config from YAML.
func GetConfig() *Config {
	var once sync.Once
	cfg := &Config{}
	once.Do(func() {
		if err := cleanenv.ReadConfig("config.yml", cfg); err != nil {
			log.Println(err)
			help, err := cleanenv.GetDescription(cfg, nil)
			if err != nil {
				log.Println(err)
			}
			log.Println(help)
		}
	})
	return cfg
}

// APP_PORT=8001 go run cmd/main.go
