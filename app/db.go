package app

import (
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (s *Server) Migrate() {
	s.db.AutoMigrate(&Todo{})
}

func (s *Server) initDB() {

	dsn := s.Cnf.String("db.conn_string")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		s.log.Fatal("Can't connect to db, ", err)
		os.Exit(1)
	}

	s.db = db

	s.Migrate()
}
