package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Host string
		Port int
	}
	Database struct {
		Addr           string
		User           string
		Pass           string
		DBName         string
		Net            string
		MaxConnections int `toml:"max_connections"`
	}
}

func LoadConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	cfg.Server.Host = os.Getenv("SERVER_HOST")
	cfg.Server.Port, _ = strconv.Atoi(os.Getenv("SERVER_PORT"))

	cfg.Database.Addr = os.Getenv("DB_ADDR")
	cfg.Database.User = os.Getenv("DB_USER")
	cfg.Database.DBName = os.Getenv("DB_NAME")
	cfg.Database.Pass = os.Getenv("DB_PASS")
	cfg.Database.Net = os.Getenv("DB_NET")
	cfg.Database.MaxConnections, _ = strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))

	if cfg.Database.Addr == "" || cfg.Database.User == "" || cfg.Database.Pass == "" {
		return nil, fmt.Errorf("missing critical database configuration (DB_USER, DB_PASSWORD, or DB_NAME)")
	}

	return cfg, nil
}
