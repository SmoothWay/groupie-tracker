package groupie

import (
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		Templates.ExecuteTemplate(w, "errors.html", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		Templates.ExecuteTemplate(w, "errors.html", http.StatusMethodNotAllowed)
		return
	}
	Templates.ExecuteTemplate(w, "mainPage.html", SearchArtist)
}

func ArtPage(w http.ResponseWriter, r *http.Request) {
	index, err := strconv.Atoi(r.RequestURI[9:])
	if err != nil {
		w.WriteHeader(400)
		Templates.ExecuteTemplate(w, "errors.html", http.StatusBadRequest)
		return
	}

	if !isValid(index) {
		w.WriteHeader(404)
		Templates.ExecuteTemplate(w, "errors.html", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		Templates.ExecuteTemplate(w, "errors.html", http.StatusMethodNotAllowed)
		return
	}

	page := &ArtOutput{}
	page.Art = SearchArtist.Artists[index-1]
	page.Rel = SearchArtist.Relations[index-1]
	Templates.ExecuteTemplate(w, "ArtPage.html", page)
}
