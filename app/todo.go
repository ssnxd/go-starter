package app

import (
	"net/http"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model

	Title       string    `db:"title"`
	Description string    `db:"description"`
	DueDate     time.Time `db:"due_date"`
}

func (s *Server) addTodo() http.HandlerFunc {

	// Creates a new todo in db
	newTodo := func(title string, description string, dueDate time.Time) (*Todo, error) {

		todo := &Todo{Title: title, Description: description, DueDate: dueDate}
		result := s.db.Create(&todo)

		if result.Error != nil {
			return nil, result.Error
		}

		return todo, nil
	}

	type Request struct {
		Title       string    `validate:"required" json:"title"`
		Description string    `validate:"required" json:"description"`

		// Due date is optional 
		DueDate     *time.Time `validate:"datetime" json:"due_date,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		var body Request

		err := s.parseBody(w, r, &body)


		if err != nil {
			s.log.Warn("[Failed parseBody read] ", err)
			s.respond(w, r, "the body is not ", http.StatusBadRequest)
			return
		}

		err = s.validate.Struct(body)

		s.log.Info(err)

		if err != nil {
			s.respond(w, r, "invalid request body", http.StatusBadRequest)
			return
		}

		todo, err := newTodo(body.Title, body.Description, *body.DueDate)

		if err != nil {
			s.log.Error("Something went wrong, try again: ", err)
			s.respond(w, r, "Something went wrong, try again", http.StatusInternalServerError)
			return
		}

		s.respond(w, r, todo, http.StatusOK)
		return
	}
}
