package apiserver

// Config ...
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogFilePath string `toml:"log_file_path"`
	LogLevel    string `toml:"log_level"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LogLevel:    "debug",
		LogFilePath: "apiserver.log",
	}
}
