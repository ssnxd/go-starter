package app

import (
	"encoding/json"
	"net/http"
	"time"

	"server/util"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/knadh/koanf"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Server struct {
	router *mux.Router
	Cnf    *koanf.Koanf
	log    *log.Logger
	db     *gorm.DB
	validate *validator.Validate
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {

	w.WriteHeader(status)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			s.log.Info("json encode: %s\n", err)
		}
	}

}

func (s *Server) logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		s.log.Info("[", r.Method, "] uri: ", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func (s *Server) onStart() {
	s.log.Info("Server listening on port ", s.Cnf.String("server.port"))
}

func New() *Server {

	router := mux.NewRouter()

	cnf := koanf.New(".")
	log := log.New()
	db := &gorm.DB{}

	validator := validator.New()

	// @TODO Implement a custom validator for time.Time
	validator.RegisterCustomTypeFunc(util.ValidateDate, time.Time{})

	srv := &Server{router, cnf, log, db, validator}

	// Load config
	srv.LoadConfig()

	// Start services
	srv.initDB()
	srv.routes()

	srv.onStart()
	return srv
}
