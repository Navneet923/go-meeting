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

	// Serve any app related static content
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("server/static")))

	// Define various application routes here
	// POST Routes:
	router.HandleFunc("/api/new_meeting", CreateNewMeeting).Methods("POST")

	// GET Routes:
	router.HanldeFunc("/api/meetings", FetchMeetings).Methods("GET")
	router.HandleFunc("/api/meeting/{id}", FetchMeetingWithId).Methods("GET")

	// Use the above router for all routes
	http.Handle("/", router)
}

// SendHTTPStatusCode is a small helper function which takes in a
// http.ResponseWriter and a status code, sets the header type and sends it
func SendHTTPStatusCode(response http.ResponseWriter, status int) {
	response.Header().Set("Content-Type", "application/json;charset=UTF-8")
	response.WriteHeader(status)
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

	// Read the request body
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	defer request.Body.Close()

	// Craft a new meeting, and handle errors
	var newMeeting NewMeeting
	if err := json.Unmarshal(body, &newMeeting); err != nil {
		// Send a mangled request status back
		SendHTTPStatusCode(response, 422)
		// TODO: Investigate what it means for the response to encode a JSON err
		if err := json.NewEncoder(response).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	// TODO: Add a newMeeting to the DB here
	SendHTTPStatusCode(response, http.StatusCreated)
}

// FetchMeetings is a GET handler which returns a list of all meetings
func FetchMeetings(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("TODO: Implement FetchMeetings"))
}

// FetchMeetingWithId is a GET handler which returns a given
// meeting (or an error) given an ID
func FetchMeetingWithId(response http.ResponseWriter, request *http.Request) {
	// Extract id
	id := mux.Vars(request)["id"]
	response.Write([]byte(fmt.Sprintf("%V id recieved.\n", id)))
}
