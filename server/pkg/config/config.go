package config

import "os"

type Config struct {
	Port         string `mapstructure:"PORT"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig() (c Config) {
	port := os.Getenv("PORT")
	jwt := os.Getenv("JWT_SECRET_KEY")

	cfg := &Config{Port: port, JWTSecretKey: jwt}
	return *cfg
}
