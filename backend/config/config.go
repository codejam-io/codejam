package config

import (
	"github.com/pelletier/go-toml/v2"
	"os"
)

type Config struct {
	Server   ServerConfig
	Database DBConfig
	Redis    RedisConfig
	OAuth    OAuthConfig
}

type ServerConfig struct {
	Listen string
}

type DBConfig struct {
	Url            string
	MaxConnections int32
}

type RedisConfig struct {
	Protocol string
	Address  string
	Password string
	Size     int
}

type OAuthConfig struct {
	Provider    string
	Id          string
	Secret      string
	RedirectUrl string
	Scopes      []string
}

func (config *Config) LoadFromFile(filename string) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = toml.Unmarshal(contents, &config)
	if err != nil {
		panic(err)
	}
}
