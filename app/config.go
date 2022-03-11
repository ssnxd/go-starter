package app

import (
	"log"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

func (s *Server) LoadConfig() {
	if err := s.Cnf.Load(file.Provider("./app/config.yml"), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
}
