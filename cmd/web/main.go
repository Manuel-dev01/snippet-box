package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Manuel-dev01/snippet-box/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
)

type contextKey string
var contextKeyUser = contextKey("user")

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	session *sessions.Session
	snippets *mysql.SnippetModel
	templateCache map[string]*template.Template
	users 	  *mysql.UserModel
}

func main() {
	// f, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer f.Close()

	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL database")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret Key")
	
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	// initialize a new template cache...
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	// Use the sessions.New() func to initialize a new session manager
	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true
	session.SameSite = http.SameSiteStrictMode

	app := &application {
		errorLog: errorLog,
		infoLog: infoLog,
		session: session,
		snippets: &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
		users: 	  &mysql.UserModel{DB: db},
	}

	// initialize a tls.Config struct to hold the non-default TLS settings we want the
	// server to use
	tlsConfig := &tls.Config {
		PreferServerCipherSuites: true,
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}
	
	srv := &http.Server {
		Addr: *addr,
		ErrorLog: errorLog,
		//Handler: mux,
		Handler: app.routes(),
		TLSConfig: tlsConfig,

		//Add Idle, Read and Write timeouts to the server
		IdleTimeout: time.Minute,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	
	infoLog.Printf("Starting server on %s", *addr)
	
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	//err = srv.ListenAndServe()
	//err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if  err = db.Ping(); err != nil {
		return nil, err
	}
	
	return db, nil
}