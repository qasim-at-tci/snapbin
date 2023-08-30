package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func uptimeHandler(w http.ResponseWriter, r *http.Request) {
	howOld := time.Now().UTC()
	fmt.Fprintf(w, "Snapbin lives (since %s)", howOld.Local().Format("2006-01-02 15:04:05"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Homepage"))
}

func snippetViewHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Snippet number %d..", id)
}

func snippetCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Snippet Create"))
}
