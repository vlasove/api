package apiserver

import "fmt"

// DBConnector ...
type DBConnector struct {
	Host         string `toml:"host"`
	Port         string `toml:"port"`
	UserName     string `toml:"username"`
	UserPassword string `toml:"user_password"`
	DBName       string `toml:"db_name"`
	SSLMode      string `toml:"sslmode"`
}

// NewDBConnector ...
func NewDBConnector() *DBConnector {
	return &DBConnector{
		Host:         "localhost",
		Port:         "5432",
		UserName:     "postgres",
		UserPassword: "postgres",
		DBName:       "postgres_test",
		SSLMode:      "disable",
	}
}

// BuildDatabaseURL ...
func BuildDatabaseURL(c *DBConnector) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host,
		c.Port,
		c.UserName,
		c.UserPassword,
		c.DBName,
		c.SSLMode,
	)
}
