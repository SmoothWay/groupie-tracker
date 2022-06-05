package groupie

import "html/template"

var Templates, TempErr = template.ParseGlob("ui/templates/*html")

const (
	PORT         = "8080"
	ArtistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	RelationsURL = "https://groupietrackers.herokuapp.com/api/relation"
	LocationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	DatesURL     = "https://groupietrackers.herokuapp.com/api/dates"
)

type ArtistData struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"fristAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Index struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type ArtOutput struct {
	A ArtistData
	R Index
}

var SearchArtist struct {
	Artists   []ArtistData
	Relations []Index `json:"index"`
}
