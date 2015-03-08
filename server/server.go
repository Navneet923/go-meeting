package server

import (
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

	// Serve any app related static content
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("server/static")))

	// Define various application routes here
	// POST Routes:
	router.HandleFunc("/api/new_meeting", CreateNewMeeting).Methods("POST")

	// GET Routes:
	router.HandleFunc("/api/meetings", FetchMeetings).Methods("GET")
	router.HandleFunc("/api/meeting/{id}", FetchMeetingWithID).Methods("GET")

	// Use the above router for all routes
	http.Handle("/", router)
}

// Render404 will render an error page
func Render404(response http.ResponseWriter, request *http.Request) {
	fmt.Printf("GET /404 Inovked")

	// Respond with a 404 page, this is a placeholder
	response.Write([]byte("Hmm looks like we 404'd trying to find: " + request.URL.Path))
}
