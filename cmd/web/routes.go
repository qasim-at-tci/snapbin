package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileserver := http.FileServer(http.Dir("./ui/static"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileserver))

	router.HandlerFunc(http.MethodGet, "/", app.homeHandler)
	router.HandlerFunc(http.MethodGet, "/snap/view/:id", app.snippetViewHandler)
	router.HandlerFunc(http.MethodGet, "/snap/create", app.snippetCreateHandler)
	router.HandlerFunc(http.MethodPost, "/snap/create", app.snippetCreatePostHandler)

	router.HandlerFunc(http.MethodGet, "/uptime", app.uptimeHandler)

	standard := alice.New(app.recoverPanic, app.logRequests, secureHeaders)

	return standard.Then(router)
}
