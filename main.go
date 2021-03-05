package main

import (
	"canitconnect/web/app"
	"net/http"
)

func main() {
	http.Handle("/", app.Index())
	// to serve css, these files are served using fileserver, with path stripped out
	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
