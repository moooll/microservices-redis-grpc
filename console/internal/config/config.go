package config

type Config struct {
	GRPCPosition string `env:"GRPC_POSITION_SERVICE"`
	GRPCPrice string `env:"GRPC_PRICE_SERVICE"`
}