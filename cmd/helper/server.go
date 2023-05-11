package helper

import (
	"net/http"
	"time"
)

type HttpConfig struct {
	ReadTimeout  int
	WriteTimeout int
	Port         string
}

func StartServer(handler http.Handler, config HttpConfig) error {
	server := http.Server{
		ReadTimeout:  time.Duration(config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
		Handler:      handler,
	}
	if config.ReadTimeout > 0 {
		server.ReadTimeout = time.Duration(config.ReadTimeout) * time.Second
	}
	if config.WriteTimeout > 0 {
		server.WriteTimeout = time.Duration(config.WriteTimeout) * time.Second
	}

	// ServeHTTP will do these things:
	// - listen on the given port, use socketmaster if exists
	// - setup signal handler which compatible with upstart & ctrl-c
	// - call graceful shutdown when the signal come.
	return ServeHTTP(&server, config.Port, 0) // use '0' for any default timeout defined from TDK
}
