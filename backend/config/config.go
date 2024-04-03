package config

import (
	"github.com/pelletier/go-toml/v2"
	"os"
)

type DBConfig struct {
	Url            string
	MaxConnections int32
}

type GitHubConfig struct {
	Id          string
	Secret      string
	RedirectUrl string
}

type ServerConfig struct {
	Listen string
}

type Config struct {
	Server   ServerConfig
	Database DBConfig
	GitHub   GitHubConfig
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
