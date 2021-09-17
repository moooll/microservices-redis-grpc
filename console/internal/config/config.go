// Package config contains configurations, such as connection strings, env variables
package config

// Config contains configuration strings
type Config struct {
	GRPCPosition string `env:"GRPC_POSITION_SERVICE"`
	GRPCPrice string `env:"GRPC_PRICE_SERVICE"`
}