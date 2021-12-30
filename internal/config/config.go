package config

import (
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Server struct {
	Port            string        `yaml:"port" env:"PORT" env-default:"8000"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" env:"SHUTDOWN_TIMEOUT" env-default:"5000000000"`
}

type ClientServer struct {
	Port string `yaml:"client_port" env:"CLIENT_PORT" env-default:"8001"`
}

type OAuth struct {
	VkClientID         string `yaml:"vk_client_id" env:"VK_CLIENT_ID" env-required:"true"`
	VkClientSecret     string `yaml:"vk_client_secret" env:"VK_CLIENT_SECRET" env-required:"true"`
	VkRedirectURL      string `yaml:"vk_redirect_url" env:"VK_REDIRECT_URL" env-required:"true"`
	GoogleClientID     string `yaml:"google_client_id" env:"GOOGLE_CLIENT_ID" env-required:"true"`
	GoogleClientSecret string `yaml:"google_client_secret" env:"GOOGLE_CLIENT_SECRET" env-required:"true"`
	GoogleRedirectURL  string `yaml:"google_redirect_url" env:"GOOGLE_REDIRECT_URL" env-required:"true"`
}

type DB struct {
	URL string `yaml:"url" env:"PG_DB_URL"`
}

type Config struct {
	Server       Server       `yaml:"server"`
	ClientServer ClientServer `yaml:"client_server"`
	OAuth        OAuth        `yaml:"oauth"`
	DB           DB           `yaml:"db"`
}

var (
	instance  *Config
	once      sync.Once
	configErr error
)

func GetConfig() (*Config, error) {
	once.Do(func() {
		instance = &Config{}
		configErr = cleanenv.ReadConfig("../../configs/config.yml", instance)
	})
	return instance, configErr
}
