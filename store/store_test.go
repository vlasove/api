package store_test

import (
	"os"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/vlasove/api/store"
)

var (
	pathToTestConfig = "../configs/teststore.toml"
	config           *store.Config
)

func TestMain(m *testing.M) {
	// Need to read data from .env configs
	config = store.NewConfig()
	_, err := toml.DecodeFile(pathToTestConfig, config)
	if err != nil {
		config.Host = "localhost"
		config.Port = "5432"
		config.UserName = "postgres"
		config.UserPassword = "postgres"
		config.DBName = "test_database" //need your local test database
		config.SSLMode = "disable"
	}
	os.Exit(m.Run())
}
