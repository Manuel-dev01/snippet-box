package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {
	f, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC)

	app := &application {
		errorLog: errorLog,
		infoLog: infoLog,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	infoLog.Printf("Starting a new Server on %s", *addr)
	
	srv := &http.Server {
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: mux,
	}
	
	infoLog.Printf("Starting server on %s", *addr)
	
	err = srv.ListenAndServe()
	//err := http.ListenAndServe(*addr, mux)

	if err != nil {
		errorLog.Fatal(err)
	}
}