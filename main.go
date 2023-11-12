package main

import (
	"database/sql"
	"log"
	api "myproject/api/library"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:K.jyothi9@2@tcp(127.0.0.1:3306)/project?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	api.RegisterRoutes(db)

	log.Println("server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
