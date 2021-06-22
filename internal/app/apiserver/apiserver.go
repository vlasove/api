package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/vlasove/api/internal/app/store/sqlstore"
)

// Start ...
func Start(config *Config) error {
	config.DatabaseURL = BuildDatabaseURL(config.DatabaseConnector)
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	store := sqlstore.New(db)
	srv := newServer(store)

	return http.ListenAndServe(config.BindAddr, srv)
}

// newDB ...
func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// // APIServer ...
// type APIServer struct {
// 	config *Config
// 	logger *logrus.Logger
// 	router *mux.Router
// 	server *http.Server
// 	store  *store.Store
// }

// // New ...
// func New(config *Config) *APIServer {
// 	config.DatabaseURL = BuildDatabaseURL(config.DatabaseConnector)
// 	return &APIServer{
// 		config: config,
// 		logger: logrus.New(),
// 		router: mux.NewRouter(),
// 	}
// }

// // Start ...
// func (s *APIServer) Start() error {
// 	if err := s.configureLogger(); err != nil {
// 		return err
// 	}
// 	s.logger.Infof("Starting API Server at port %s", s.config.BindAddr)

// 	s.configureRouter()
// 	s.logger.Info("Routes configurete successfully")
// 	s.server = &http.Server{
// 		Addr:    s.config.BindAddr,
// 		Handler: s.router,
// 	}

// 	if err := s.configureStore(); err != nil {
// 		return err
// 	}
// 	dbname, username := s.store.DatabaseInfo()
// 	s.logger.Infof("Database connection successfully opened. DB:%s USER:%s", dbname, username)
// 	return s.server.ListenAndServe()
// }

// // Shutdown ...
// func (s *APIServer) Shutdown(ctx context.Context) error {
// 	return s.server.Shutdown(ctx)
// }

// // Logger ...
// func (s *APIServer) Logger() *logrus.Logger {
// 	return s.logger
// }

// // configureStore ...
// func (s *APIServer) configureStore() error {
// 	st := store.New(s.config.Store)
// 	if err := st.Open(); err != nil {
// 		return err
// 	}
// 	s.store = st
// 	return nil
// }

// // Store ...
// func (s *APIServer) Store() *store.Store {
// 	return s.store
// }

// // configureLogger ...
// func (s *APIServer) configureLogger() error {
// 	level, err := logrus.ParseLevel(s.config.LogLevel)
// 	if err != nil {
// 		return err
// 	}
// 	s.logger.SetLevel(level)

// 	dir, err := filepath.Abs(filepath.Dir(s.config.LogFilePath))
// 	if err != nil {
// 		return err
// 	}
// 	if _, err := os.Stat(dir); os.IsNotExist(err) {
// 		err = os.Mkdir(dir, os.ModePerm)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	file, err := os.OpenFile(s.config.LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 	if err != nil {
// 		return err
// 	}
// 	s.logger.SetOutput(io.MultiWriter(file, os.Stdout))
// 	return nil
// }

// // configureRouter ...
// func (s *APIServer) configureRouter() {
// 	s.router.HandleFunc("/hello", s.handleHello())
// }

// // handleHello ...
// func (s *APIServer) handleHello() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		s.logger.Infof("Method:%s URL:%s User-Agent:%s", r.Method, r.URL.String(), r.UserAgent())
// 		_, err := io.WriteString(w, "Hello, world!")
// 		if err != nil {
// 			s.logger.Errorf(
// 				"error occured while writing response to Method:%s URL:%s Error:%s",
// 				r.Method,
// 				r.URL.String(),
// 				err.Error(),
// 			)
// 		}
// 	}
// }
