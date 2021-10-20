package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/veremchukvv/stonks-test/pkg/logging"
	"sync"
)

type Server struct {
	Port string `yaml:"port" env:"PORT"`
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

func GetConfig() *Config {
	once.Do(func(){
		logger := logging.NewLogger(false, "console")
		logger.Info("read application configuration")
		instance = &Config{}
		err := cleanenv.ReadConfig("../config.yml", instance)
		if err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}