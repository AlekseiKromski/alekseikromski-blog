package api

import (
	"alekseikromski.com/blog/api/storage"
	v1 "alekseikromski.com/blog/api/v1"
	_ "alekseikromski.com/blog/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

type Server struct {
	config *Config
	mux    *http.ServeMux
	apis   []Api
}

func NewServer(config *Config, storage storage.Storage) *Server {
	return &Server{
		mux:    http.NewServeMux(),
		config: config,
		apis: []Api{
			v1.NewV1(storage),
		},
	}
}

func (s *Server) Start() error {
	log.Println("Register handlers")

	fs := http.FileServer(http.Dir("./front-end/build/"))
	s.mux.Handle("/", fs)

	log.Println("Route [ / ] was mounted - GENERIC - FRONT-END")

	s.mux.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("200 OK"))
	})

	log.Println("Route [ /swagger ] was mounted - GENERIC")
	s.mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	log.Println("Route [ /healthz ] was mounted - GENERIC")

	for _, api := range s.apis {
		api.RegisterRoutes()
		api.Mount(s.mux)
	}

	log.Printf("Run server on %s", s.config.addr)
	err := http.ListenAndServe(s.config.addr, s.mux)
	return err
}
