package groupie

import (
	"log"
	"net/http"
	"strconv"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		log.Println("Status: Not Found (404)")
		Templates.ExecuteTemplate(w, "error.html", 404)
		return
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		log.Printf("Status: 405. %v Method is Not Allowed", r.Method)
		Templates.ExecuteTemplate(w, "error.html", http.StatusMethodNotAllowed)
		return
	}
	Templates.ExecuteTemplate(w, "index.html", SearchArtist)
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.RequestURI[9:])

	if id < 1 || id > 52 {
		w.WriteHeader(404)
		Templates.ExecuteTemplate(w, "error.html", http.StatusNotFound)
		log.Println("Status: 404 Page not Found")
		return
	}
	if err != nil {
		w.WriteHeader(400)
		Templates.ExecuteTemplate(w, "error.html", http.StatusBadRequest)
		log.Println("Status: Bad Request (400)")
		return
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		Templates.ExecuteTemplate(w, "error.html", http.StatusMethodNotAllowed)
		log.Printf("Status: 405. %v Method is Not Allowed", r.Method)
		return
	}
	res := &ArtOutput{}
	res.A = SearchArtist.Artists[id-1]
	res.R = SearchArtist.Relations[id-1]
	Templates.ExecuteTemplate(w, "artist.html", res)
}
