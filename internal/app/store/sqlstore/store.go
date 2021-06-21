package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq" // Anon import for pq driver
	"github.com/vlasove/api/internal/app/store"
)

// Store ...
type Store struct {
	//config         *Config 			TO DELETE
	db             *sql.DB
	userRepository *UserRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// DatabaseURL ...
// func (s *Store) DatabaseURL() string {
// 	return fmt.Sprintf(
// 		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
// 		s.config.Host,
// 		s.config.Port,
// 		s.config.UserName,
// 		s.config.UserPassword,
// 		s.config.DBName,
// 		s.config.SSLMode,
// 	)
// }

// DatabaseInfo ...
// func (s *Store) DatabaseInfo() (dbname, user string) {
// 	dbname = s.config.DBName
// 	user = s.config.UserName
// 	return
// }

// User returned public repo interface
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}

// Open ...
// func (s *Store) Open() error {
// 	db, err := sql.Open("postgres", s.DatabaseURL())
// 	if err != nil {
// 		return err
// 	}
// 	if err := db.Ping(); err != nil {
// 		return err
// 	}
// 	s.db = db
// 	return nil
// }

// // Close ...
// func (s *Store) Close() {
// 	s.db.Close()
// }
