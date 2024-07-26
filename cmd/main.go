package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"main.go/Errors"
	"main.go/Fetch"
)

func server(w http.ResponseWriter, r *http.Request) {
	// Use a switch statement to handle different URL paths
	jsonArtistsCards := Fetch.Fetch_cards(w, r)

	switch r.URL.Path {
	case "/":
		// Construct the path to the index.html template file
		IndexPath := filepath.Join("..", "templates", "index.html")

		// Parse the index.html template
		IndexParse, err := template.ParseFiles(IndexPath)
		if err != nil {
			// Call a custom error handler function for HTTP 500 errors
			Errors.Error500(w, r)
		}

		// Execute the index.html template
		err = IndexParse.ExecuteTemplate(w, "index.html", jsonArtistsCards)
		if err != nil {
			Errors.Error500(w, r)
			return
		}
	case "/profile":
		Profile_Path := filepath.Join("..", "templates", "profile.html")
		// Parse the index.html template
		Profile_Parse, err := template.ParseFiles(Profile_Path)

		var artist_info interface{}
		artist_info = Fetch.Fetch_profile(w, r, jsonArtistsCards)

		Profile_Parse.ExecuteTemplate(w, "profile.html", artist_info)
		if err != nil {
			// Call a custom error handler function for HTTP 500 errors
			Errors.Error500(w, r)
		}

	default:
		// Handle requests for paths other than "/"
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404")
	}
}

func main() {
	// Handle requests to the root URL ("/")
	http.HandleFunc("/", server)
	styles := http.FileServer(http.Dir("../stylesheets"))
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", styles))
	// Start the HTTP server on port 1234
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println("Error in starting the server", err)
	}
}
