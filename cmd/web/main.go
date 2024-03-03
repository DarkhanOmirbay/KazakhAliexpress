package main

import (
	"KazakhAliexpress/SE2220/pkg/models/postgresql"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Application struct {
	DB        *sql.DB
	ItemModel *postgresql.ItemModel
	//Items []*models.Item
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	connstring := "user=postgres dbname=KazakhAliexpress password='703905' host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		log.Fatal(err)
	}
	app := Application{
		DB: db,
	}
	server := http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}
	err = server.ListenAndServe()
	fmt.Printf("Starting server on %s", *addr)
	log.Fatal(err)
}
