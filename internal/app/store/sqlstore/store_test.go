package sqlstore_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/BurntSushi/toml"
)

var (
	pathToTestConfig = "../configs/teststore.toml"
	connStr          string
)

// Connector ...
type Connector struct {
	Host         string `toml:"host"`
	Port         string `toml:"port"`
	UserName     string `toml:"username"`
	UserPassword string `toml:"user_password"`
	DBName       string `toml:"db_name"`
	SSLMode      string `toml:"sslmode"`
}

// NewConnector ...
func NewConnector() *Connector {
	return &Connector{
		Host:         "localhost",
		Port:         "5432",
		UserName:     "postgres",
		UserPassword: "postgres",
		DBName:       "postgres_test",
		SSLMode:      "disable",
	}
}

// DatabaseURL ...
func (c *Connector) DatabaseURL() string {
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

func TestMain(m *testing.M) {
	// Need to read data from .env configs
	conn := NewConnector()
	_, err := toml.DecodeFile(pathToTestConfig, conn)
	if err != nil {
		conn.Host = "localhost"
		conn.Port = "5432"
		conn.UserName = "postgres"
		conn.UserPassword = "postgres"
		conn.DBName = "test_database" //need your local test database
		conn.SSLMode = "disable"
	}
	connStr = conn.DatabaseURL()
	os.Exit(m.Run())
}
