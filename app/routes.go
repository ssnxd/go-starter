package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	POST string = "POST"
	GET  string = "GET"
)

func (s *Server) parseBody(w http.ResponseWriter, r *http.Request, ref interface{}) error {

	// should acceptes body from request and parse it into struct
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &ref)
}

func (s *Server) Handle(path string, f func(http.ResponseWriter,
	*http.Request), method string) *mux.Route {

	s.log.Info("[", method, "] uri: ", path, " registered")

	return s.router.HandleFunc(path, f).Methods(method)
}

func (s *Server) routes() {

	// POST todo /todo {Creates a new todo}
	s.Handle("/todo", s.addTodo(), POST)

	// Init global middlewares
	s.router.Use(s.logMiddleware)
}
