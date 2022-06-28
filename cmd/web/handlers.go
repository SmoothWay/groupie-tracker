package main

import (
	"net/http"
	"strconv"

	groupie "github.com/SmoothWay/groupie-tracker/pkg"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	app.render(w, r, "home-page.html", &templateData{Data: groupie.SearchArtist})
}

func (app *application) artPage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	app.infoLog.Println(id)
	if err != nil || !isValid(id) {
		app.notFound(w)
		return
	}

	if r.Method != http.MethodGet {
		app.clientError(w, 405)
	}

	page := &groupie.ArtOutput{}
	page.Art = groupie.SearchArtist.Artists[id-1]
	page.Rel = groupie.SearchArtist.Relations[id-1]
	app.render(w, r, "art-page.html", &templateData{Page: page})
}
