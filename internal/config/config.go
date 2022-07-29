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
	Host     string `yaml:"host" env:"HOSTPG" env-default:"localhost"`
	Port     string `yaml:"port" env:"PORTPG" env-default:"5432"`
	Username string `yaml:"username" env:"USERNAME" env-default:"pg"`
	Password string `yaml:"password" env:"PASSWORD" env-default:"pass"`
	DBName   string `yaml:"dbName" env:"DBNAME" env-default:"crud"`
	SSLMode  string `yaml:"sslMode" env:"SSLMODE" env-default:"disable"`
}

type App struct {
	Port string `yaml:"port" env:"PORTAPP" env-default:"8000"`
	//  PORTAPP=8001 go run cmd/main.go
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
