package main

import (
	"alekseikromski.com/blog/api"
	"alekseikromski.com/blog/api/guard"
	"alekseikromski.com/blog/api/guard/jwt"
	"alekseikromski.com/blog/api/storage/dbstore"
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

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	jwtSecret := os.Getenv("jwtSecret")
	googCaptchaToken := os.Getenv("GOOGLE_CAPTCHA_TOKEN")

	if username == "" || password == "" || hostname == "" || port == "" || database == "" {
		log.Fatalf("[ERROR] Database credits is required")
	}
	config := api.NewConfig(addr)

	//Prepare storage object
	dbstore, err := dbstore.NewDbConnection(
		username,
		password,
		hostname,
		port,
		database,
	)

	if err != nil {
		log.Fatalf("There is the toruble with db connection: %v", err)
		return
	}

	//Register guards
	guards := []guard.Guard{
		jwt.New(jwtSecret),
	}

	//Create router
	router := router.NewRouter(guards)

	//Prepare apis
	apis := []api.Api{
		v1.NewV1(dbstore, router, googCaptchaToken, guards),
	}

	server := api.NewServer(config, router, apis)
	log.Println("[INFO] Create server instance")
	err = server.Start()
	if err != nil {
		log.Fatalf("server error: %v", err)
	}

	log.Println("Server down")
}
