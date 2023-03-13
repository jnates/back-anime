package main

import (
	"backend_crudgo/infrastructure"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	log.Println("Starting API CMD")
	port := os.Getenv("API_PORT")
	infrastructure.Start(port)
}
