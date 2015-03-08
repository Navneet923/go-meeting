package server

import "net/http"

// SendHTTPStatusCode is a small helper function which takes in a
// http.ResponseWriter and a status code, sets the header type and sends it
func SendHTTPStatusCode(response http.ResponseWriter, status int) {
	response.Header().Set("Content-Type", "application/json;charset=UTF-8")
	response.WriteHeader(status)
}
