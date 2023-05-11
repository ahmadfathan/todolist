package config

import (
	"github.com/ahmadfathan/todolist/cmd/helper"
)

type AppConfig struct {
	Server Server
}

type Server struct {
	HTTP helper.HttpConfig
}

func GetAppConfig() AppConfig {

	httpConfig := helper.HttpConfig{
		ReadTimeout:  8,
		WriteTimeout: 8,
		Port:         ":3030",
	}

	server := Server{
		HTTP: httpConfig,
	}

	cfg := AppConfig{
		Server: server,
	}

	return cfg
}
