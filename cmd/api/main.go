package main

import (
	"database/sql"
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

	//db, err := app.openDB()
	app.check(err)

	port := os.Getenv("PORT")
	portStr := fmt.Sprintf(":%s", port)

	log.Printf("Starting server on port %s", portStr)

	err = http.ListenAndServe(portStr, app.routes())
	app.check(err)
}

func (app *application) getConnectionString() string {
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	// CONFIG: add pw if your DB user has a password
	//	pw := os.Getenv("PG_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s "+
		" sslmode=disable", host, port, user, dbName)

	return connStr
}

func (app *application) openDB() (*sql.DB, error) {
	connStr := app.getConnectionString()

	db, err := sql.Open("postgres", connStr)
	app.check(err)

	err = db.Ping()
	app.check(err)

	return db, nil
}
