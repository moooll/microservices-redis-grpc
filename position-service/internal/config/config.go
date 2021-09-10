package config

type Config struct {
	GRPCAddr string `env:"GRPC_URI"`
}
