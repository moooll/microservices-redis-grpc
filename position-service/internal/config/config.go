package config

type Config struct {
	GRPCClientAddr string `env:"GRPC_CLIENT_URI"`
	GRPCServerPort string `env:"GRPC_SERVER_PORT"`
}
