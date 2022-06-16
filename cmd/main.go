package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	groupie "timshowtime.net/groupie-tracker/pkg"
)

func main() {
	port := os.Getenv("PORT")
	err1 := groupie.Unmarshal(groupie.UrlArt, &groupie.SearchArtist.Artists)
	err2 := groupie.Unmarshal(groupie.UrlRel, &groupie.SearchArtist)
	if err1 != nil || err2 != nil {
		fmt.Println("ERROR")
		return
	}

	if groupie.TempErr != nil {
		log.Fatal(groupie.TempErr)
	}
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))
	mux.HandleFunc("/", groupie.Home)
	mux.HandleFunc("/artists/", groupie.ArtPage)
	fmt.Printf("Starting server - http://localhost:%v\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
