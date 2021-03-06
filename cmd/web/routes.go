package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/artists", app.artPage)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return app.recoverPanic(app.logRequest(secureHeaders((mux))))
}
