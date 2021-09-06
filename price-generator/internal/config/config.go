package config

type Config struct {
	RedisURI string `env:"REDIS_URI"`
}