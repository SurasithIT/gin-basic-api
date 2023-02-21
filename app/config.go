package app

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTPPort   int    `envconfig:"HTTP_PORT" default:"8080"`
	HTTPPrefix string `envconfig:"HTTP_PATH_PREFIX" default:""`
}

var AppConfig = &Config{}

func LoadConfig() {
	err := envconfig.Process("", AppConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
}
