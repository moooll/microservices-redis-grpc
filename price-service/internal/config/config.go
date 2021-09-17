// Package config contains configurations, such as connection strings, env variables
package config

// Config contains configuration strings
type Config struct {
	RedisURI string `env:"REDIS_URI"`
	PortGRPC string `env:"PORT_GRPC"`
}