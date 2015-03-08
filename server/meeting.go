package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

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
	fmt.Printf("GET /meetings Invoked")

	// Respond with a list of meetings here
	response.Write([]byte("TODO: Implement FetchMeetings"))
}

// FetchMeetingWithID is a GET handler which returns a given
// meeting (or an error) given an ID
func FetchMeetingWithID(response http.ResponseWriter, request *http.Request) {
	// Extract id
	id := mux.Vars(request)["id"]

	fmt.Printf("GET /meeting/%s\n", id)

	// Respond with meeting specific data
	response.Write([]byte(fmt.Sprintf("%V id recieved.\n", id)))
}
