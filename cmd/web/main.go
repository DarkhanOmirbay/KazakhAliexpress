package main

import (
	"KazakhAliexpress/SE2220/pkg/models/postgresql"
	"database/sql"
	"flag"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorlog *log.Logger
	infolog  *log.Logger
	items    *postgresql.ItemModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// info log
	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// err log
	errorlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	connstring := "user=postgres dbname=KazakhAliexpress password='703905' host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		log.Fatal(err)
	}
	app := &application{
		errorlog: errorlog,
		infolog:  infolog,
		items:    &postgresql.ItemModel{DB: db},
	}
	server := http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: errorlog,
	}
	err = server.ListenAndServe()
	infolog.Printf("Starting server on %s", *addr)
	errorlog.Fatal(err)
}
