package pkg

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost         string
	PostgresPort         string
	PostgresUser         string
	PostgresDatabase     string
	PostgresPassword     string
	DatabaseUrl          string
	Environment          string
	LogLevel             string
	HTTP_PORT            string
	PostServiceHost      string
	PostServicePort      string
	CollectorServiceHost string
	CollectorServicePort string
}

func Load() Config {
	c := Config{}

	c.PostgresHost = cast.ToString(GetOrReturnDefault("POSTGRES_HOST", "db"))
	c.PostgresPort = cast.ToString(GetOrReturnDefault("POSTGRES_PORT", "5432"))
	c.PostgresUser = cast.ToString(GetOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(GetOrReturnDefault("POSTGRES_PASSWORD", "postgres"))
	c.PostgresDatabase = cast.ToString(GetOrReturnDefault("POSTGRES_DATABSE", "data_db"))
	c.DatabaseUrl = cast.ToString(GetOrReturnDefault("DATABASE_URL", "db"))
	c.Environment = cast.ToString(GetOrReturnDefault("ENVIRONMENT", "info"))
	c.LogLevel = cast.ToString(GetOrReturnDefault("LOG_LEVEL", "debug"))
	c.CollectorServiceHost = cast.ToString(GetOrReturnDefault("COLLECTOR_SERVICE_HOST", "localhost"))
	c.CollectorServicePort = cast.ToString(GetOrReturnDefault("COLLECTOR_SERVICE_PORT", ":5000"))
	c.PostServiceHost = cast.ToString(GetOrReturnDefault("POST_SERVICE_HOST", "localhost"))
	c.PostServicePort = cast.ToString(GetOrReturnDefault("POST_SERVICE_PORT", ":5001"))
	c.HTTP_PORT = cast.ToString(GetOrReturnDefault("HTTP_PORT", ":5002"))

	return c
}

func GetOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
