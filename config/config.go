package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	App string

	Environment string // development, staging, production

	LogLevel string // debug, info, warn, error, dpanic, panic, fatal

	ServiceScheme string
	ServiceHost   string
	ServicePort   string

	HTTPHost string
	HTTPPort string

	BasePath string
	//services
	TodoServiceURL string

	//queryparam configs
	DefaultOffset string
	DefaultLimit  string

	//postgres
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string

	//rabbit
	RabbitURI string
}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.App = cast.ToString(getOrReturnDefaultValue("APP", "todo_api-gateway"))

	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "development"))

	config.LogLevel = cast.ToString(getOrReturnDefaultValue("LOG_LEVEL", "debug"))

	//Service configs
	config.ServiceScheme = cast.ToString(getOrReturnDefaultValue("SERVICE_SCHEME", "http"))
	config.ServiceHost = cast.ToString(getOrReturnDefaultValue("SERVICE_HOST", "localhost"))
	config.ServicePort = cast.ToString(getOrReturnDefaultValue("SERVICE_PORT", ":8001"))
	//HTTP Configs
	config.HTTPHost = cast.ToString(getOrReturnDefaultValue("HTTP_HOST", "localhost"))
	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8001"))

	config.BasePath = cast.ToString(getOrReturnDefaultValue("BASE_PATH", "/v1"))
	//QueryParamconfigs
	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))

	//Services
	config.TodoServiceURL = cast.ToString(getOrReturnDefaultValue("COURIER_SERVICE_URL", "http://localhost:8000"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "postgres"))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "123"))

	config.RabbitURI = cast.ToString(getOrReturnDefaultValue("AMQP_URI", "amqp://guest:guest@localhost:5672"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
