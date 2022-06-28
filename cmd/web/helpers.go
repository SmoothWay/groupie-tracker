package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"time"
)

func Unmarshal(s string, a interface{}) error {
	dataArt, err := http.Get(s)
	if err != nil {
		return err
	}
	defer dataArt.Body.Close()
	body, err := ioutil.ReadAll(dataArt.Body)
	if err != nil {
		return err
	}
	err2 := json.Unmarshal(body, a)
	if err2 != nil {
		return err
	}
	return nil
}

func isValid(n int) bool {
	if n < 1 || n > 52 {
		return false
	}
	return true
}

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.CurrentYear = time.Now().Year()
	return td
}
func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("the template %s does not exits", name))
		return
	}
	buf := new(bytes.Buffer)

	err := ts.Execute(buf, app.addDefaultData(td, r))
	if err != nil {
		app.serverError(w, err)
		return
	}
	buf.WriteTo(w)
}

func (app *application) execTemp(w http.ResponseWriter, status int) {
	templates, templErr := template.ParseFiles("./ui/templates/errors.html")
	if templErr != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		app.errorLog.Fatal(templErr)
	}
	templates.Execute(w, status)
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)
	w.WriteHeader(http.StatusInternalServerError)
	app.execTemp(w, http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	app.execTemp(w, status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
