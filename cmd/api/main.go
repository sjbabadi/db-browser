package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type application struct{}

func main() {
	app := &application{}

	err := godotenv.Load()
	app.check(err)

	port := os.Getenv("PORT")
	portStr := fmt.Sprintf(":%s", port)

	log.Printf("Starting server on port %s", portStr)

	err = http.ListenAndServe(portStr, app.routes())
	app.check(err)
}
