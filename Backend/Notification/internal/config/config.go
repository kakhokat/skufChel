package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type Config struct {
	HttpServer HttpServer  `json:"http_server"`
	Database   Database    `json:"database"`
	Kafka      KafkaConfig `json:"kafka"`
	Mail       MailAccount `json:"mail"`
	Salt       string      `json:"salt"`
}

type HttpServer struct {
	Port          string `json:"port" env:"SERVER_PORT"`
	TimeoutSecond int    `json:"timeout_second" env:"SERVER_TIMEOUT"`
}

type Database struct {
	Hostname string `json:"host" env:"DB_HOSTNAME"`
	Port     int    `json:"port" env:"DB_PORT"`
	Username string `json:"username" env:"DB_USERNAME"`
	Password string `json:"password" env:"DB_PASSWORD"`
	Database string `json:"database" env:"DB_DATABASE"`
	Sslmode  string `json:"sslmode" env:"DB_SSLMODE"`
}

type KafkaConfig struct {
	HostName  string `json:"host" env:"KAFKA_HOSTNAME"`
	Port      string `json:"port" env:"KAFKA_PORT"`
	Topic     string `json:"topic" env:"KAFKA_TOPIC"`
	Partition int    `json:"partition" env:"KAFKA_PARTITION"`
} //todo env добавить

type MailAccount struct {
	Mail     string `json:"email" env:"SENDER_MAIL"`
	Password string `json:"password" env:"SENDER_PASSWORD"`
} //todo тут тоже env

func CreateConfig() (*Config, error) {
	configPath := "config/config.json"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.Errorf("config file %s does not exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, errors.Wrap(err, "cannot read config")
	}

	_ = cleanenv.ReadEnv(&cfg)

	//todo validator для конфига нужно сделать, чтобы поля обязательные точно заполнялись
	return &cfg, nil
}
