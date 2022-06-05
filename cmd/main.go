package main

import (
	groupie "groupie/pkg"
	"log"
)

func main() {
	err := groupie.JsonUnmarshal(&groupie.SearchArtist.Artists)
	if err != nil {
		log.Fatal(err)
	}
	groupie.Handlers()
}
