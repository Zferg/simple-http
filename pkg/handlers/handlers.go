package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// ServeFiles accepts a path and fetches it within the
// static directory
func ServeFiles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Serve a static asset
	http.ServeFile(w, r, fmt.Sprintf("/static/%s", vars["filename"]))
	log.Println("Sending the gophers to fetch", vars["filename"])

	// Handle error if no file is found
	_, err := w.Write([]byte("file not find"))
	if err != nil {
		log.Println("failed to write response", err)
		return
	}
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
	log.Println("Sending the gophers to fetch", r.RequestURI)
}
