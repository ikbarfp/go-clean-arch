package server

import (
	"log"

	"github.com/ikbarfp/go-clean-arch/pkg/config"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (s *Server) Run() {

	viperConf := config.New(config.Config{
		Name:      "config",
		Extension: "yml",
		Path:      ".",
	})

	err := viperConf.Load()
	if err != nil {
		log.Fatalf("[ERROR] failed to load configuration file : %v", err)
	}

}
