package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func (app *application) uptimeHandler(w http.ResponseWriter, r *http.Request) {
	howOld := time.Now().UTC()
	fmt.Fprintf(w, "Snapbin lives - %s", howOld.Local().Format("2006-01-02 15:04:05"))
}

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) snippetViewHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Snippet number %d..", id)
}

func (app *application) snippetCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)

		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Snippet Create"))
}
