package main

import (
	"KazakhAliexpress/SE2220/pkg/models/postgresql"
	"database/sql"
	"flag"
	"github.com/golangcollege/sessions"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	errorlog      *log.Logger
	infolog       *log.Logger
	items         *postgresql.ItemModel
	templateCache map[string]*template.Template
	session       *sessions.Session
	users         *postgresql.UserModel
	carts         *postgresql.CartModel
	orders        *postgresql.OrderModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	flag.Parse()

	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errorlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	connstring := "user=postgres dbname=KazakhAliexpress password='703905' host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		log.Fatal(err)
	}
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorlog.Fatal(err)
	}
	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true
	session.SameSite = http.SameSiteStrictMode

	app := &application{
		errorlog:      errorlog,
		infolog:       infolog,
		items:         &postgresql.ItemModel{DB: db},
		templateCache: templateCache,
		session:       session,
		users:         &postgresql.UserModel{DB: db},
		carts:         &postgresql.CartModel{DB: db},
		orders:        &postgresql.OrderModel{DB: db},
	}
	server := http.Server{
		Addr:         *addr,
		Handler:      app.routes(),
		ErrorLog:     errorlog,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err = server.ListenAndServe()
	//err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	infolog.Printf("Starting server on %s", *addr)
	errorlog.Fatal(err)
}
