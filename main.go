package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// (1) Define a home handler function which write a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the wildcard from the request using r.PathValue()
	// and try to convert it to an integer using the strconv.Atoi() function. If
	// it can't be converted to an integer, or the value is less than 1, we
	// return a 404 page not found response.
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
	// phased out after refactoring in line above ^ -- w.Write([]byte("Display a specific snippet..."))
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Save a new snippet..."))
}

func main() {
	// (2) Use the http.NewServeMux() function to initialize a new servemux, then
	// (3) register the home function as the handler for "/" URL pattern.
	// ** by avoiding http.DefaultServeMux as global variable and using our own locally-scoped
	// servemux, we are writing more clear, maintainable and secure code as
	// third-party packages will be unable to register routes; a best practice
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	// Register the two new handler functions and corressponding route patterns with
	// the servemux, in exactly the same way we did before.
	mux.HandleFunc("GET /snippet/view/{id}", snippetView) // Add the {id} wildcard segment
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	// Create the new route, which is restricted to POST requests only.
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// Print a log message to say that the server is starting.
	log.Print("starting server on :4000")

	// Use the http.ListenAndServe() function to start a new web server. We passs in
	// two parameters: The TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() as always non-nil.

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
