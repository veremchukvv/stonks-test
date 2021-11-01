package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
	"time"
)

type Server struct {
	Port string `yaml:"port" env:"PORT" env-default:"8000"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" env:"SHUTDOWN_TIMEOUT" env-default:"5000000000"`
}

type ClientServer struct {
	Port string `yaml:"client_port" env:"CLIENT_PORT" env-default:"8001"`
}

type OAuth struct {
	VkClientID string `yaml:"vk_client_id" env:"VK_CLIENT_ID" env-required:"true"`
	VkClientSecret string `yaml:"vk_client_secret" env:"VK_CLIENT_SECRET" env-required:"true"`
	VkRedirectURL string `yaml:"vk_redirect_url" env:"VK_REDIRECT_URL" env-required:"true"`
	}

type DB struct {

}

type Config struct {
	Server Server `yaml:"server"`
	ClientServer ClientServer `yaml:"client_server"`
	OAuth OAuth `yaml:"oauth"`
	DB DB
}

var instance *Config
var once sync.Once
var configErr error

func GetConfig() (*Config, error) {
	once.Do(func(){
		instance = &Config{}
		configErr = cleanenv.ReadConfig("../../configs/config.yml", instance)
	})
	return instance, configErr
}