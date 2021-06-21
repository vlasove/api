package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Anon import for pq driver
)

// Store ...
type Store struct {
	config *Config
	db     *sql.DB
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// DatabaseURL ...
func (s *Store) DatabaseURL() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		s.config.Host,
		s.config.Port,
		s.config.UserName,
		s.config.UserPassword,
		s.config.DBName,
		s.config.SSLMode,
	)
}

// DatabaseInfo ...
func (s *Store) DatabaseInfo() (dbname, user string) {
	dbname = s.config.DBName
	user = s.config.UserName
	return
}

// Open ...
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.DatabaseURL())
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}
