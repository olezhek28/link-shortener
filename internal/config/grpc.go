package config

import (
	"errors"
	"os"
)

const (
	grpcHostEnvName = "GRPC_PORT"
)

type GRPCConfig interface {
	Host() string
}

type grpcConfig struct {
	host string
}

// GetGRPCConfig ...
func GetGRPCConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	return &grpcConfig{
		host: host,
	}, nil
}

// Host ...
func (cfg *grpcConfig) Host() string {
	return cfg.host
}
