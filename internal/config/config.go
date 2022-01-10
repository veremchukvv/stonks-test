package config

import (
	"os"
	"sync"
	"time"

	"github.com/veremchukvv/stonks-test/pkg/logging"

	"github.com/ilyakaznacheev/cleanenv"
)

type Server struct {
	Port            string        `yaml:"port" env:"PORT" env-default:"8000"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" env:"SHUTDOWN_TIMEOUT" env-default:"5000000000"`
	CORS            []string      `yaml:"cors" env:"CORS" env-separator:"," env-default:"http://127.0.0.1:3000 http://localhost:3000"`
}

type Client struct {
	Port           string `yaml:"client_port" env:"CLIENT_PORT" env-default:"8001"`
	ReactClientURL string `yaml:"react_client_url" env:"REACT_URL" env-default:"http://localhost:3000"`
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
	URL string `yaml:"url" env:"DATABASE_URL"`
}

type Config struct {
	Server Server `yaml:"server"`
	Client Client `yaml:"client_server"`
	OAuth  OAuth  `yaml:"oauth"`
	DB     DB     `yaml:"db"`
}

var (
	instance  *Config
	once      sync.Once
	configErr error
)

func GetConfig() (*Config, error) {
	log := logging.NewLogger(false, "console")

	once.Do(func() {
		e := os.Getenv("IS_PRODUCTION")
		if e == "" {
			instance = &Config{}
			configErr = cleanenv.ReadConfig("../../configs/config.yml", instance)
			log.Info("Loading config from /configs/config.yml")
		} else {
			instance = &Config{}
			configErr = cleanenv.ReadConfig("./configs/config_example.yml", instance)
			log.Info("Loading config from /configs/config_example.yml")
		}
	})
	return instance, configErr
}
