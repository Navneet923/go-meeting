package main 

import (
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
)

// Setup the routes for the server and any static assets
func SetupServer() {
    router := mux.NewRouter()

    // Define error'd route handler
    router.NotFoundHandler = http.HandlerFunc(Render404)

    // Define various application routes here
    router.HandleFunc("/", RenderIndex).Methods("GET")
        
    // Serve any app related static content
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("go-meeting/static"))))

    // Use the above router for all routes
    http.Handle("/", router)
}

// Render 404 page
func Render404(response http.ResponseWriter, request *http.Request) {
    response.Write([]byte("Hmm looks like we 404'd trying to find: " + request.URL.Path))
}

// Render Home page
func RenderIndex(response http.ResponseWriter, request *http.Request) {
    response.Write([]byte("Hello world!"))
}

// Main entry point
func main() {
    // Setup the server and get things going
    SetupServer()

    fmt.Printf("Server up and running....\n")
    http.ListenAndServe(":3000", nil)
}
