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
	err := jsonUnmarshal(ArtistsURL, &SearchArtist.Artists)
	if err != nil {
		log.Fatal(err)
	}

	err = jsonUnmarshal(RelationsURL, &SearchArtist)
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static/"))))
	mux.HandleFunc("/", HomePage)
	mux.HandleFunc("/artists/", ArtistPage)
	fmt.Printf("Starting server - http://localhost:%v\n", PORT)
	if err := http.ListenAndServe(":"+PORT, mux); err != nil { // start the server
		log.Fatal(err)
	}
}

func jsonUnmarshal(URL string, a interface{}) error {
	resp, err := http.Get(URL)
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
