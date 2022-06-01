package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

var cfg Config

type Config struct {
	PgPort     string `yaml:"pgPort" env:"pgPort" env-default:"5432"`
	PgHost     string `yaml:"pgHost" env:"pgHost" env-default:"localhost"`
	DBType     string `yaml:"dbType" env:"dbType" env-default:"postgres"`
	PgUser     string `yaml:"pgUser" env:"pgUser" env-default:"pg"`
	PgPassword string `yaml:"pgPassword" env:"pgPassword" env-default:"pass"`
	DB         string `yaml:"db" env:"db" env-default:"crud"`
	Migrate    bool   `yaml:"migrate" env:"migrate" env-default:"false"`
	LogLevel   string `yaml:"logLevel" env:"logLevel" env-default:"info"`
	PORT       int    `yaml:"port" env:"port" env-default:"4001"`
	HOST       string `yaml:"host" env:"port" env-default:"localhost"`
}

func InitConfig() error {
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return err
	}
	return nil
}

func SetConfig(newConfig Config) {
	cfg = newConfig
}

func GetConfig() Config {
	return cfg
}
