package store

// Config ...
type Config struct {
	Host         string `toml:"host"`
	Port         string `toml:"port"`
	UserName     string `toml:"username"`
	UserPassword string `toml:"user_password"`
	DBName       string `toml:"db_name"`
	SSLMode      string `toml:"sslmode"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Host:         "localhost",
		Port:         "5432",
		UserName:     "postgres",
		UserPassword: "postgres",
		DBName:       "postgres",
		SSLMode:      "disable",
	}
}
