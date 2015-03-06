package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRoutes defines the various routes that this server will handle
// as well as allow for the hosting of static assets
func SetupRoutes() {
	router := mux.NewRouter()

	// Define error'd route handler
	router.NotFoundHandler = http.HandlerFunc(Render404)

	// Define various application routes here
	router.HandleFunc("/api/new_meeting", CreateNewMeeting).Methods("POST")

	// Serve any app related static content
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("server/static")))

	// Use the above router for all routes
	http.Handle("/", router)
}

// Render404 will render an error page
func Render404(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hmm looks like we 404'd trying to find: " + request.URL.Path))
}

type tests struct {
	Test string
}

// CreateNewMeeting is a POST handler for a request to make a new meeting
func CreateNewMeeting(response http.ResponseWriter, request *http.Request) {
	fmt.Printf("POST /api/new_meeting called\n")

	decoder := json.NewDecoder(request.Body)
	var t tests
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("BODY: %v\n", t)
}
