package config

import (
	"EM-TASK/app/internal/infrastructure/logger"
	"github.com/ilyakaznacheev/cleanenv"
)

var Config struct {
	HTTP struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"http"`
	Logger             logger.Config `yaml:"logger"`
	DatabaseConnString string        `yaml:"database_conn_string"`
}

func Init(configPath string) error {
	err := cleanenv.ReadConfig(configPath, &Config)
	return err
}
