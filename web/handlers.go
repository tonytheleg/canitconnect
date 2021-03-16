package web

import (
	v1 "canitconnect/internal/api/v1"
	"net/http"
)

// CallCurl handles the route for curling and endpoint
func CallCurl() http.Handler {
	return http.HandlerFunc(v1.Curl)
}

// CallTraceroute handles the route for performing a traceroute to an endpoint
func CallTraceroute() http.Handler {
	return http.HandlerFunc(v1.Traceroute)
}

// CallNetcat handles the route for performing a netcat to test a port
func CallNetcat() http.Handler {
	return http.HandlerFunc(v1.Netcat)
}
