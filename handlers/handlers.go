package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler servers the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "<h1>Welcome to the Home Page</h1>")
}

// HelloHandler serves a simple hello message.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "<h1>Hello, world!</h1>")
}