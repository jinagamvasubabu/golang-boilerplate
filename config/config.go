package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

var cfg Config

type Config struct {
	PgPort             string `yaml:"pgPort" env:"pgPort" env-default:"5432"`
	PgHost             string `yaml:"pgHost" env:"pgHost" env-default:"localhost"`
	DBType             string `yaml:"dbType" env:"dbType" env-default:"postgres"`
	PgUser             string `yaml:"pgUser" env:"pgUser" env-default:"pg"`
	PgPassword         string `yaml:"pgPassword" env:"pgPassword" env-default:"pass"`
	MongoPort          string `yaml:"mongoPort" env:"mongoPort" env-default:"27017"`
	MongoHost          string `yaml:"mongoHost" env:"mongoHost" env-default:"localhost"`
	DB                 string `yaml:"db" env:"db" env-default:"crud"`
	Collection         string `yaml:"collection" env:"collection" env-default:"crudCollection"`
	Migrate            bool   `yaml:"migrate" env:"migrate" env-default:"true"`
	LogLevel           string `yaml:"logLevel" env:"logLevel" env-default:"info"`
	PORT               int    `yaml:"port" env:"port" env-default:"4001"`
	HOST               string `yaml:"host" env:"port" env-default:"localhost"`
	MaxOpenConnections int    `yaml:"maxOpenConnections" env:"port" env-default:"5"`
	KafkaBrokerUrl     string `yaml:"kafkaBrokerUrl" env:"kafkaBrokerUrl" env-default:"localhost:29092"`
	Topic              string `yaml:"topic" env:"topic" env-default:"events.BookCreated"`
}

func InitConfig() error {
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return err
	}
	return nil
}

func SetConfig(newConfig Config, DBType string) {
	cfg = newConfig
	cfg.DBType = DBType
}

func GetConfig() Config {
	return cfg
}
