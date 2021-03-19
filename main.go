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
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	// API Routes
	http.Handle("/api/v1/curl", web.CallCurl())
	http.Handle("/api/v1/traceroute", web.CallTraceroute())
	http.Handle("/api/v1/netcat", web.CallNetcat())

	// App Form Routes
	http.Handle("/curl", web.CallCurl())
	http.Handle("/traceroute", web.CallTraceroute())
	http.Handle("/netcat", web.CallNetcat())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
