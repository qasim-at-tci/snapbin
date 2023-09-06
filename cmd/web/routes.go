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

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.homeHandler))
	router.Handler(http.MethodGet, "/snap/view/:id", dynamic.ThenFunc(app.snippetViewHandler))
	router.Handler(http.MethodGet, "/snap/create", dynamic.ThenFunc(app.snippetCreateHandler))
	router.Handler(http.MethodPost, "/snap/create", dynamic.ThenFunc(app.snippetCreatePostHandler))

	router.HandlerFunc(http.MethodGet, "/uptime", app.uptimeHandler)

	standard := alice.New(app.recoverPanic, app.logRequests, secureHeaders)

	return standard.Then(router)
}
