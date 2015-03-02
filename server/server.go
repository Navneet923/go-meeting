package server 

import (
    "github.com/gorilla/mux"
    "net/http"
    "fmt"
)

// Setup the routes for the server and any static assets
func SetupRoutes() {
    router := mux.NewRouter()

    // Define error'd route handler
    router.NotFoundHandler = http.HandlerFunc(Render404)

    // Define various application routes here
    router.HandleFunc("/api/add/{a:[0-9]+}/{b:[0-9]+}", AddHandler).Methods("GET")
        
    // Serve any app related static content
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("server/static")))

    // Use the above router for all routes
    http.Handle("/", router)
}

// Render 404 page
func Render404(response http.ResponseWriter, request *http.Request) {
    response.Write([]byte("Hmm looks like we 404'd trying to find: " + request.URL.Path))
}

// Add Handler (POC entrypoint - just adds two numbers)
func AddHandler(response http.ResponseWriter, request *http.Request) {
    fmt.Printf("Add handler invoked: %v\n", mux.Vars(request))
    vars := mux.Vars(request)
    a, b := vars["a"], vars["b"]
    response.Write([]byte(a + b))
}

