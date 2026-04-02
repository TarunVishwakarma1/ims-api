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
	return &Config{
		Application: Application{
			Port: getEnv("APPLICATION_PORT", "8082"),
		},
		WebApp: WebApp{
			Port: getEnv("WEB_APP_PORT", "ims"),
			Name: getEnv("WEB_APP_NAME", "3000"),
		},
	}
}

func getEnv(key string, dValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return dValue
}
