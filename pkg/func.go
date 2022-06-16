package groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Unmarshal(s string, a interface{}) error {
	dataArt, err1 := http.Get(s)
	if err1 != nil {
		fmt.Println("No  response from request")
	}
	defer dataArt.Body.Close()
	body, err := ioutil.ReadAll(dataArt.Body) // response body is []byte
	err2 := json.Unmarshal(body, a)
	if err2 != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	return err
}

func isValid(n int) bool {
	if n < 1 || n > 52 {
		return false
	}
	return true
}
