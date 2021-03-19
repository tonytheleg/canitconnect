package main

import (
	"canitconnect/web"
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("web/templates/*.html"))
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func main() {
	http.HandleFunc("/", serveIndex)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/api/v1/curl", web.CallCurl())
	http.Handle("/api/v1/traceroute", web.CallTraceroute())
	http.Handle("/api/v1/netcat", web.CallNetcat())
	http.Handle("/curl", web.CallCurlForm())
	http.Handle("/traceroute", web.CallTraceroute())
	http.Handle("/netcat", web.CallNetcat())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
