package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/uptime", app.uptimeHandler)
	mux.HandleFunc("/", app.homeHandler)
	mux.HandleFunc("/snippet/view", app.snippetViewHandler)
	mux.HandleFunc("/snippet/create", app.snippetCreateHandler)

	return mux
}
