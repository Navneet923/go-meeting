package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// NewMeeting defines a structure which contains the various parameters
// which come to us from the client
type NewMeeting struct {
	Title     string `json:"title"`
	Attendees string `json:"attendees"`
	Notes     string `json:"notes"`
}

// CreateNewMeeting is a POST handler for a request to make a new meeting
func CreateNewMeeting(response http.ResponseWriter, request *http.Request) {
	fmt.Printf("POST /api/new_meeting called\n")

	var newMeeting NewMeeting
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	if err := request.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &newMeeting); err != nil {
		response.Header().Set("Content-Type", "application/json;charset=UTF-8")
		response.WriteHeader(422)
		if err := json.NewEncoder(response).Encode(err); err != nil {
			panic(err)
		}
	}

	fmt.Printf("TODO: Add meeting to DB\n")
	response.Header().Set("Content-Type", "application/json;charset=UTF-8")
	response.WriteHeader(http.StatusCreated)
	fmt.Printf("Got New Meeting request: %V\n", newMeeting)
}
