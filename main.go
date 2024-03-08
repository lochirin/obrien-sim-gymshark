package main

import (
	"example/gymshark/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	router := routes.PackageRoutes()
	http.Handle("/api", router)
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
