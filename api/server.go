package api

import (
	_ "alekseikromski.com/blog/docs"
	router "alekseikromski.com/blog/router"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
)

type Server struct {
	config *Config
	apis   []Api
	router *router.Router
}

func NewServer(config *Config, router *router.Router, apis []Api) *Server {
	return &Server{
		config: config,
		router: router,
		apis:   apis,
	}
}

func (s *Server) Start() error {
	log.Println("[INFO] Register handlers")

	s.router.CreateRoute(
		"/",
		http.MethodGet,
		func(writer http.ResponseWriter, request *http.Request) {
			file, _ := os.ReadFile("./front-end/build/index.html")
			writer.Write(file)
		},
		nil,
	)

	s.router.CreateRoute(
		"/static/*",
		http.MethodGet,
		func(writer http.ResponseWriter, request *http.Request) {
			fileServer := http.FileServer(http.Dir("./front-end/build/"))
			fileServer.ServeHTTP(writer, request)
		},
		nil,
	)

	s.router.CreateRoute(
		"/healthz",
		http.MethodGet,
		func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte("200 OK"))
		},
		nil,
	)
	s.router.CreateRoute(
		"/swagger/*",
		http.MethodGet,
		httpSwagger.WrapHandler,
		nil,
	)

	for _, api := range s.apis {
		if err := api.RegisterRoutes(); err != nil {
			log.Fatalf("there is the problem with routes registration: %v", err)
		}
	}

	log.Printf("Run server on %s", s.config.addr)
	err := http.ListenAndServe(s.config.addr, s.router)
	return err
}
