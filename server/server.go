package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Main entry point
func main() {
	router := mux.NewRouter()

	// Define error'd route handler
	router.NotFoundHandler = http.HandlerFunc(Render404)

	// Define various application routes here
	router.HandleFunc("/", RenderIndex).Methods("GET")

	// Use the above router for all routes
	http.Handle("/", router)

	fmt.Printf("Server up and listening...")
	http.ListenAndServe("0.0.0.0:3000", nil)
}

// Render 4040 page
func Render404(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hmm looks like we 404'd trying to find: " + request.URL.Path))
}

// Render Home page
func RenderIndex(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hello world!"))
}
