package handlers

import (
	"log"
	"net/http"
)

// TODO:
func LogRequestHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		// call the original http.Handler we're wrapping
		h.ServeHTTP(w, r)

		// gather information about request and log it
		uri := r.URL.String()
		method := r.Method

		log.Println(method, uri)
	}

	// http.HandlerFunc wraps a function so that it
	// implements http.Handler interface
	return http.HandlerFunc(fn)
}
