package config

import "os"

type Config struct {
	Hostname     string `mapstructure:"HOSTNAME"`
	Server       string `mapstructure:"PORT"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig() (c Config) {
	server := os.Getenv("GRPC_SERVER")
	jwt := os.Getenv("JWT_SECRET_KEY")
	hostname := os.Getenv("HOSTNAME")

	cfg := &Config{Server: server, JWTSecretKey: jwt, Hostname: hostname}
	return *cfg
}
