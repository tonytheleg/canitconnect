package main

import (
	"canitconnect/web"
	"net/http"
)

func main() {
	http.Handle("/", web.Index())
	http.Handle("/api/v1/curl", web.CallCurl())
	http.Handle("/api/v1/traceroute", web.CallTraceroute())
	http.Handle("/api/v1/netcat", web.CallNetcat())
	// to serve css, these files are served using fileserver, with path stripped out
	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
