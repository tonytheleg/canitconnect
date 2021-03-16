package main

import (
	"canitconnect/web"
	"fmt"
	"net/http"
)

func main() {
	//http.Handle("/", web.Index())
	http.Handle("/api/v1/curl", web.CallCurl())
	http.Handle("/api/v1/traceroute", web.CallTraceroute())
	http.Handle("/api/v1/netcat", web.CallNetcat())

	err := http.ListenAndServe(":8080", http.FileServer(http.Dir("web/assets")))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
