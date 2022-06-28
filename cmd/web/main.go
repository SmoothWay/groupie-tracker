package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	groupie "github.com/SmoothWay/groupie-tracker/pkg"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	templateCache map[string]*template.Template
}

func main() {

	addr := flag.String("addr", ":8080", "HTTP network address")

	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	tempalteCache, err := newTemplateCache("./ui/templates/")
	if err != nil {
		errorLog.Fatal(err)
	}
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		templateCache: tempalteCache,
	}

	infoLog.Printf("Starting server on %s\n", *addr)
	err1 := Unmarshal(groupie.UrlArt, &groupie.SearchArtist.Artists)
	err2 := Unmarshal(groupie.UrlRel, &groupie.SearchArtist)
	if err1 != nil || err2 != nil {
		errorLog.Fatal("Cannot Unmarshal JSON")
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
