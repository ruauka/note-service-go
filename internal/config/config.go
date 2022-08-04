package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Postgres `yaml:"postgres"`
	App      `yaml:"app"`
}

type Postgres struct {
	Host     string `yaml:"host" env:"PSQL_HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"PSQL_PORT" env-default:"5432"`
	Username string `yaml:"username" env:"PSQL_USERNAME" env-default:"pg"`
	Password string `yaml:"password" env:"PSQL_PASSWORD" env-default:"pass"`
	DBName   string `yaml:"dbName" env:"PSQL_DBNAME" env-default:"crud"`
	SSLMode  string `yaml:"sslMode" env:"PSQL_SSLMODE" env-default:"disable"`
}

type App struct {
	Port string `yaml:"port" env:"APP_PORT" env-default:"8000"`
	//  APP_PORT=8001 go run cmd/main.go
}

func GetConfig() *Config {
	var once sync.Once
	cfg := &Config{}
	once.Do(func() {
		if err := cleanenv.ReadConfig("config.yml", cfg); err != nil {
			help, _ := cleanenv.GetDescription(cfg, nil)
			log.Println(help)
			log.Println(err)
		}
	})
	return cfg
}
