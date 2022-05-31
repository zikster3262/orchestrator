package config

import "os"

type Config struct {
	Port string `mapstructure:"PORT"`
}

func LoadConfig() (c Config) {
	port := os.Getenv("PORT")

	cfg := &Config{Port: port}
	return *cfg
}
