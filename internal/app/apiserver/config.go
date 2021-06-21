package apiserver

import "github.com/vlasove/api/store"

// Config ...
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogFilePath string `toml:"log_file_path"`
	LogLevel    string `toml:"log_level"`
	Store       *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LogLevel:    "debug",
		LogFilePath: "apiserver.log",
		Store:       store.NewConfig(),
	}
}
