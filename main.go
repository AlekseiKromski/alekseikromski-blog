package main

import (
	"alekseikromski.space/api"
	"log"
	"os"
)

func main() {

	addr := os.Getenv("BLOG_ADDRESS")
	if len(addr) == 0 {
		log.Fatalf("env BLOG_ADDRESS is required")
	}

	config := api.NewConfig(addr)
	server := api.NewServer(config)
	log.Println("Create server instance")
	err := server.Start()
	if err != nil {
		log.Fatalf("server error: %v", err)
	}

	log.Println("Server down")
}
