// Package config contains configurations strings for network connections
package config

// Config describes configuration of the service
type Config struct {
	GRPCClientAddr string `env:"GRPC_CLIENT_URI"`
	GRPCServerPort string `env:"GRPC_SERVER_PORT"`
	PgURI          string `env:"PG_URI"`
	ServerID       int    `env:"SERVER_ID"`
}
