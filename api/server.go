package api

import (
	v1 "alekseikromski.space/api/v1"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Server struct {
	config *Config
	mux    *http.ServeMux
	apis   []Api
}

func NewServer(config *Config) *Server {
	return &Server{
		mux:    http.NewServeMux(),
		config: config,
		apis: []Api{
			v1.NewV1(),
		},
	}
}

func (s *Server) Start() error {
	log.Println("Register handlers")

	s.mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		filePath := filepath.Join("front-end", "index.html")
		file, err := os.ReadFile(filePath)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf("cannot find front-end: %v", err)))
			return
		}

		writer.Write(file)
	})
	log.Println("Route [ / ] was mounted - GENERIC - FRONT-END")

	s.mux.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("200 OK"))
	})

	log.Println("Route [ /healthz ] was mounted - GENERIC")

	for _, api := range s.apis {
		api.RegisterRoutes()
		api.Mount(s.mux)
	}

	log.Printf("Run server on %s", s.config.addr)
	err := http.ListenAndServe(s.config.addr, s.mux)
	return err
}
