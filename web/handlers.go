package web

import (
	v1 "canitconnect/internal/api/v1"
	"fmt"
	"net/http"
)

// Index returns the root page
func Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "CanItConnect?!\n")
	})
}

// CallCurl handles the route for curling and endpoint
func CallCurl() http.Handler {
	return http.HandlerFunc(v1.ShowCurlData)
}

// CallTraceroute handles the route for performing a traceroute to an endpoint
func CallTraceroute() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Traceroute Called!\n")
	})
}

// CallNetcat handles the route for performing a netcat to test a port
func CallNetcat() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Netcat Called!\n")
	})
}

// CallPing handles the route to performing a ping test to an endpoint
func CallPing() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ping Called!\n")
	})
}
