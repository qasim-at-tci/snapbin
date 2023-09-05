package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/qasim-at-tci/snapbin/internal/models"
)

func (app *application) uptimeHandler(w http.ResponseWriter, r *http.Request) {
	howOld := time.Now().UTC()
	fmt.Fprintf(w, "Snapbin lives - %s", howOld.Local().Format("2006-01-02 15:04:05"))
}

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	snaps, err := app.snaps.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snaps = snaps

	app.render(w, http.StatusOK, "home.tmpl", data)
}

func (app *application) snippetViewHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snap, err := app.snaps.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Snap = snap

	app.render(w, http.StatusOK, "view.tmpl", data)
}

func (app *application) snippetCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the form for creating a new snippet..."))
}

func (app *application) snippetCreatePostHandler(w http.ResponseWriter, r *http.Request) {
	title := "O Snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snaps.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
