package groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Handlers() {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static/"))))
	mux.HandleFunc("/", HomePage)
	mux.HandleFunc("/artists/", ArtistPage)
	// http.HandleFunc("/articles", )
	fmt.Printf("Starting server - http://localhost:%v\n", PORT)
	if err := http.ListenAndServe(":"+PORT, mux); err != nil { // start the server
		log.Fatal(err)
	}
}

func JsonUnmarshal(a interface{}) error {
	resp, err := http.Get(ArtistsURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, a); err != nil {
		return err
	}
	return nil
}
