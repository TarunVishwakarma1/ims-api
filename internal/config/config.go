package config

import "os"

type Config struct {
	Application
	WebApp
}

type Application struct {
	Port string
}

type WebApp struct {
	Port string
	Name string
}

func Load() *Config {
	return &Config{}
}

func getEnv(key string, dValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return dValue
}
