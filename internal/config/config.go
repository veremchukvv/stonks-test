package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
	"time"
)

type Server struct {
	Port string `yaml:"port" env:"PORT"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" env:"SHUTDOWN_TIMEOUT"`
}

type ClientServer struct {
	Port string `yaml:"client_port" env:"CLIENT_PORT"`
}

type DB struct {

}

type Config struct {
	Server Server `yaml:"server"`
	ClientServer ClientServer `yaml:"client_server"`
	DB DB
}

var instance *Config
var once sync.Once
var configErr error

func GetConfig() (*Config, error) {
	once.Do(func(){
		instance = &Config{}
		configErr = cleanenv.ReadConfig("../../configs/main.yml", instance)
	})
	return instance, configErr
}