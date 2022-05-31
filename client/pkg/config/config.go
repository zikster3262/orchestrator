package config

import "os"

type Config struct {
	Server       string `mapstructure:"PORT"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig() (c Config) {
	server := os.Getenv("GRPC_SERVER")
	jwt := os.Getenv("JWT_SECRET_KEY")

	cfg := &Config{Server: server, JWTSecretKey: jwt}
	return *cfg
}
