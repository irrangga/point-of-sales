package models

type ServerConfig struct {
	Environment string `env:"ENVIRONMENT"`
	AppPort     string `env:"APP_PORT"`
	DBUsername  string `env:"DB_USERNAME"`
	DBPassword  string `env:"DB_PASSWORD"`
	DBHost      string `env:"DB_HOST"`
	DBPort      string `env:"DB_PORT"`
	DBName      string `env:"DB_NAME"`
}
