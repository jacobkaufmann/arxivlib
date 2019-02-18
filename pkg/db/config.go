package db

import (
	"errors"
	"os"
)

const (
	HostKey     = "host"
	PortKey     = "port"
	UsernameKey = "username"
	PasswordKey = "password"
	DBKey       = "db"
)

// A Config holds configuration information about a datastore
type Config struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
}

// ConfigFromEnv returns a Config built from the specified environment variables
func ConfigFromEnv(varNames map[string]string) *Config {
	cfg := &Config{}

	cfg.Host = os.Getenv(varNames[HostKey])
	cfg.Port = os.Getenv(varNames[PortKey])
	cfg.Username = os.Getenv(varNames[UsernameKey])
	cfg.Password = os.Getenv(varNames[PasswordKey])
	cfg.DatabaseName = os.Getenv(varNames[DBKey])

	return cfg
}

// ErrEnvVarNotFound is a failure to retrieve an environment variable
var ErrEnvVarNotFound = errors.New("environment variable not found")
