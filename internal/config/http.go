package config

import (
	"errors"
	"os"
)

const (
	httpHostEnvName = "HTTP_PORT"
)

type HTTPConfig interface {
	Host() string
}

type httpConfig struct {
	host string
}

// GetHTTPConfig ...
func GetHTTPConfig() (HTTPConfig, error) {
	host := os.Getenv(httpHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("http host not found")
	}

	return &httpConfig{
		host: host,
	}, nil
}

// Host ...
func (cfg *httpConfig) Host() string {
	return cfg.host
}
