package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresURL string
	ServerPort  int
	SecretJWT   string
}

func NewParsedConfig() *Config {
	godotenv.Load()
	config := Config{}
	config.ServerPort, _ = strconv.Atoi(os.Getenv("SERVER_PORT"))
	config.SecretJWT = os.Getenv("SECRET_JWT")
	config.PostgresURL = os.Getenv("DATABASE_URL")
	// config.PostgresURL = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"))
	return &config
}
