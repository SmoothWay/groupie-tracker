package groupie

import (
	"log"
	"net/http"
	"path"
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
	log.Println(path.Base(r.RequestURI))
}
