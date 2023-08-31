package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/uptime", uptimeHandler)
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/snippet/view", snippetViewHandler)
	mux.HandleFunc("/snippet/create", snippetCreateHandler)

	log.Print("Starting server on port 4001..")
	err := http.ListenAndServe(":4001", mux)

	if err != nil {
		log.Fatal(err)
	}
}
