package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("ERROR:%s", err)
	}

	http.HandleFunc("/", WebApp)
	http.ListenAndServe(os.Getenv("PORT"), nil)
}
