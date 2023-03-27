package main

import (
	"alekseikromski.com/blog/api"
	"alekseikromski.com/blog/api/storage/memstore"
	v1 "alekseikromski.com/blog/api/v1"
	router "alekseikromski.com/blog/router"
	"log"
	"os"
)

// @title			Swagger Aleksei Kromski blog API
// @version		1.0
// @description	This is a simple api for aleksei kromski blog
// @host			localhost:3001
// @BasePath		/api/v1
func main() {

	addr := os.Getenv("BLOG_ADDRESS")
	if len(addr) == 0 {
		log.Fatalf("env BLOG_ADDRESS is required")
	}

	config := api.NewConfig(addr)

	//Prepare storage object
	memstore := memstore.NewMemStorage()

	//Create router
	router := router.NewRouter()

	//Prepare apis
	apis := []api.Api{
		v1.NewV1(memstore, router),
	}

	server := api.NewServer(config, router, apis)
	log.Println("Create server instance")
	err := server.Start()
	if err != nil {
		log.Fatalf("server error: %v", err)
	}

	log.Println("Server down")
}
