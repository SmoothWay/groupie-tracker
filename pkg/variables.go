package groupie

import "html/template"

const (
	PORT   = "8080"
	UrlArt = "https://groupietrackers.herokuapp.com/api/artists"
	UrlRel = "https://groupietrackers.herokuapp.com/api/relation"
)

var Templates, TempErr = template.ParseGlob("./ui/templates/*.html")

type DataArt struct {
	ID             int                 `json:"id"`
	Image          string              `json:"image"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	CreationDate   int                 `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	Locations      string              `json:"locations"`
	DatesLocations map[string][]string `json:"datesLocations"`

	ConcertDates string `json:"concertDates"`
}

type Index struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type ArtOutput struct {
	Art DataArt
	Rel Index
}

var SearchArtist struct {
	Artists   []DataArt
	Relations []Index `json:"index"`
}
